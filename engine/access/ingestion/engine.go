// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package ingestion

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/rs/zerolog"

	"github.com/dapperlabs/flow-go/consensus/hotstuff/model"
	"github.com/dapperlabs/flow-go/engine"
	"github.com/dapperlabs/flow-go/engine/access/rpc"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/model/messages"
	"github.com/dapperlabs/flow-go/module"
	"github.com/dapperlabs/flow-go/network"
	"github.com/dapperlabs/flow-go/state/protocol"
	"github.com/dapperlabs/flow-go/storage"
	"github.com/dapperlabs/flow-go/utils/logging"
)

// Engine represents the ingestion engine, used to funnel data from other nodes
// to a centralized location that can be queried by a user
type Engine struct {
	unit  *engine.Unit   // used to manage concurrency & shutdown
	log   zerolog.Logger // used to log relevant actions with context
	state protocol.State // used to access the  protocol state
	me    module.Local   // used to access local node information

	// Conduits
	collectionConduit network.Conduit

	// storage
	// FIX: remove direct DB access by substituting indexer module
	blocks       storage.Blocks
	headers      storage.Headers
	collections  storage.Collections
	transactions storage.Transactions

	rpcEngine *rpc.Engine
}

// New creates a new access ingestion engine
func New(log zerolog.Logger,
	net module.Network,
	state protocol.State,
	me module.Local,
	blocks storage.Blocks,
	headers storage.Headers,
	collections storage.Collections,
	transactions storage.Transactions,
	rpcEngine *rpc.Engine,
) (*Engine, error) {

	// initialize the propagation engine with its dependencies
	eng := &Engine{
		unit:         engine.NewUnit(),
		log:          log.With().Str("engine", "ingestion").Logger(),
		state:        state,
		me:           me,
		blocks:       blocks,
		headers:      headers,
		collections:  collections,
		transactions: transactions,
		rpcEngine:    rpcEngine,
	}

	collConduit, err := net.Register(engine.CollectionProvider, eng)
	if err != nil {
		return nil, fmt.Errorf("could not register collection provider engine: %w", err)
	}

	eng.collectionConduit = collConduit

	return eng, nil
}

// Ready returns a ready channel that is closed once the engine has fully
// started. For the ingestion engine, we consider the engine up and running
// upon initialization.
func (e *Engine) Ready() <-chan struct{} {
	return e.unit.Ready()
}

// Done returns a done channel that is closed once the engine has fully stopped.
// For the ingestion engine, it only waits for all submit goroutines to end.
func (e *Engine) Done() <-chan struct{} {
	return e.unit.Done()
}

// SubmitLocal submits an event originating on the local node.
func (e *Engine) SubmitLocal(event interface{}) {
	e.Submit(e.me.NodeID(), event)
}

// Submit submits the given event from the node with the given origin ID
// for processing in a non-blocking manner. It returns instantly and logs
// a potential processing error internally when done.
func (e *Engine) Submit(originID flow.Identifier, event interface{}) {
	e.unit.Launch(func() {
		err := e.process(originID, event)
		if err != nil {
			engine.LogError(e.log, err)
		}
	})
}

// ProcessLocal processes an event originating on the local node.
func (e *Engine) ProcessLocal(event interface{}) error {
	return e.Process(e.me.NodeID(), event)
}

// Process processes the given event from the node with the given origin ID in
// a blocking manner. It returns the potential processing error when done.
func (e *Engine) Process(originID flow.Identifier, event interface{}) error {
	return e.unit.Do(func() error {
		return e.process(originID, event)
	})
}

// process processes the given ingestion engine event. Events that are given
// to this function originate within the expulsion engine on the node with the
// given origin ID.
func (e *Engine) process(originID flow.Identifier, event interface{}) error {
	switch entity := event.(type) {
	case *messages.CollectionResponse:
		return e.handleCollectionResponse(originID, entity)
	default:
		return fmt.Errorf("invalid event type (%T)", event)
	}
}

// OnFinalizedBlock is called by the follower engine after a block has been finalized and the state has been updated
func (e *Engine) OnFinalizedBlock(hb *model.Block) {
	e.unit.Launch(func() {
		id := hb.BlockID
		err := e.processFinalizedBlock(id)
		if err != nil {
			e.log.Error().Err(err).Hex("block_id", id[:]).Msg("failed to process block")
			return
		}
	})
}

// processBlock handles an incoming finalized block.
func (e *Engine) processFinalizedBlock(id flow.Identifier) error {

	block, err := e.blocks.ByID(id)
	if err != nil {
		return fmt.Errorf("failed to lookup block: %w", err)
	}

	// Notify rpc handler of new finalized block height
	e.rpcEngine.SubmitLocal(block)

	// FIX: we can't index guarantees here, as we might have more than one block
	// with the same collection as long as it is not finalized

	// TODO: substitute an indexer module as layer between engine and storage

	// index the block storage with each of the collection guarantee
	err = e.blocks.IndexBlockForCollections(block.Header.ID(), flow.GetIDs(block.Payload.Guarantees))
	if err != nil {
		return fmt.Errorf("could not index block for collections: %w", err)
	}

	// request each of the collections from the collection node
	return e.requestCollections(block.Payload.Guarantees...)
}

// handleCollectionResponse handles the response of the a collection request made earlier when a block was received
func (e *Engine) handleCollectionResponse(originID flow.Identifier, response *messages.CollectionResponse) error {
	collection := response.Collection
	light := collection.Light()

	// FIX: we can't index guarantees here, as we might have more than one block
	// with the same collection as long as it is not finalized

	// store the light collection (collection minus the transaction body - those are stored separately)
	// and add transaction ids as index
	err := e.collections.StoreLightAndIndexByTransaction(&light)
	if err != nil {
		// ignore collection if already seen
		if errors.Is(err, storage.ErrAlreadyExists) {
			e.log.Debug().
				Hex("collection_id", logging.ID(light.ID())).
				Msg("collection is already seen")
			return nil
		}
		return err
	}

	// now store each of the transaction body
	for _, tx := range collection.Transactions {
		err := e.transactions.Store(tx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) requestCollections(guarantees ...*flow.CollectionGuarantee) error {
	for _, guarantee := range guarantees {
		req := &messages.CollectionRequest{
			ID:    guarantee.ID(),
			Nonce: rand.Uint64(),
		}
		err := e.collectionConduit.Submit(req, guarantee.SignerIDs...)
		if err != nil {
			return err
		}
	}

	return nil

}

// OnBlockIncorporated is a noop for this engine since access node is only dealing with finalized blocks
func (e *Engine) OnBlockIncorporated(*model.Block) {
}

// OnDoubleProposeDetected is a noop for this engine since access node is only dealing with finalized blocks
func (e *Engine) OnDoubleProposeDetected(*model.Block, *model.Block) {
}
