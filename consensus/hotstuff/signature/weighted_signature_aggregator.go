package signature

import (
	"fmt"
	"sync"

	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/crypto"
	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/signature"
)

// signerInfo holds information about a signer, its stake and index
type signerInfo struct {
	weight uint64
	index  int
}

// WeightedSignatureAggregator implements consensus/hotstuff.WeightedSignatureAggregator.
// It is a wrapper around signature.SignatureAggregatorSameMessage, which implements a
// mapping from node IDs (as used by HotStuff) to index-based addressing of authorized
// signers (as used by SignatureAggregatorSameMessage).
type WeightedSignatureAggregator struct {
	aggregator   *signature.SignatureAggregatorSameMessage // low level crypto BLS aggregator, agnostic of weights and flow IDs
	ids          flow.IdentityList                         // all possible ids (only gets updated by constructor)
	idToInfo     map[flow.Identifier]signerInfo            // auxiliary map to lookup signer weight and index by ID (only gets updated by constructor)
	totalWeight  uint64                                    // weight collected (gets updated)
	collectedIDs map[flow.Identifier]struct{}              // map of collected IDs (gets updated)
	lock         sync.RWMutex                              // lock for atomic updates to totalWeight and collectedIDs
}

var _ hotstuff.WeightedSignatureAggregator = (*WeightedSignatureAggregator)(nil)

// NewWeightedSignatureAggregator returns a weighted aggregator initialized with a list of flow
// identities, their respective public keys, a message and a domain separation tag. The identities
// represent the list of all possible signers.
//
// The constructor errors if:
// - the list of identities is empty
// - if the length of keys does not match the length of identities
// - if one of the keys is not a valid public key.
//
// A weighted aggregator is used for one aggregation only. A new instance should be used for each
// signature aggregation task in the protocol.
func NewWeightedSignatureAggregator(
	ids flow.IdentityList, // list of all authorized signers
	pks []crypto.PublicKey, // list of corresponding public keys used for signature verifications
	message []byte, // message to get an aggregated signature for
	dsTag string, // domain separation tag used by the signature
) (*WeightedSignatureAggregator, error) {
	if len(ids) != len(pks) {
		return nil, fmt.Errorf("keys length %d and identities length %d do not match", len(pks), len(ids))
	}

	// build the internal map for a faster look-up
	idToInfo := make(map[flow.Identifier]signerInfo)
	for i, id := range ids {
		idToInfo[id.NodeID] = signerInfo{
			weight: id.Stake,
			index:  i,
		}
	}

	// instantiate low-level crypto aggregator, that works based on signer indices instead of nodeIDs
	agg, err := signature.NewSignatureAggregatorSameMessage(message, dsTag, pks)
	if err != nil {
		return nil, fmt.Errorf("instantiating index-based signature aggregator failed: %w", err)
	}

	return &WeightedSignatureAggregator{
		aggregator:   agg,
		ids:          ids,
		idToInfo:     idToInfo,
		collectedIDs: make(map[flow.Identifier]struct{}),
	}, nil
}

// Verify verifies the signature under the stored public keys and message.
// Error returns:
//  - model.ErrInvalidSigner if signerID is invalid (not a consensus participant)
//  - model.ErrInvalidSignature if signerID is valid but signature is cryptographically invalid
//  - generic error in case of unexpected runtime failures
// The function is thread-safe.
func (w *WeightedSignatureAggregator) Verify(signerID flow.Identifier, sig crypto.Signature) error {
	info, ok := w.idToInfo[signerID]
	if !ok {
		return fmt.Errorf("couldn't find signerID %s in the index map: %w", signerID, model.ErrInvalidSigner)
	}

	ok, err := w.aggregator.Verify(info.index, sig) // no error expected during normal operation
	if err != nil {
		return fmt.Errorf("couldn't verify signature from %s: %w", signerID, err)
	}
	if !ok {
		return fmt.Errorf("invalid signature from %s: %w", signerID, model.ErrInvalidSignature)
	}
	return nil
}

// hasSignature returns true if the input ID already included a signature
// and false otherwise.
// The function is thread safe.
func (w *WeightedSignatureAggregator) hasSignature(signerID flow.Identifier) bool {
	w.lock.RLock()
	defer w.lock.RUnlock()
	_, ok := w.collectedIDs[signerID]
	return ok
}

// TrustedAdd adds a signature to the internal set of signatures and adds the signer's
// weight to the total collected weight, iff the signature is _not_ a duplicate.
//
// The total weight of all collected signatures (excluding duplicates) is returned regardless
// of any returned error.
// The function errors
//  - engine.InvalidInputError if signerID is invalid (not a consensus participant)
//  - engine.DuplicatedEntryError if the signer has been already added
// The function is thread-safe.
func (w *WeightedSignatureAggregator) TrustedAdd(signerID flow.Identifier, sig crypto.Signature) (uint64, error) {
	info, found := w.idToInfo[signerID]
	if !found {
		return w.TotalWeight(), engine.NewInvalidInputErrorf("couldn't find signerID %s in the map", signerID)
	}

	if w.hasSignature(signerID) {
		return w.TotalWeight(), engine.NewDuplicatedEntryErrorf("SigneID %s was already added", signerID)
	}

	// atomically update the signatures pool and the total weight
	w.lock.Lock()
	defer w.lock.Unlock()

	err := w.aggregator.TrustedAdd(info.index, sig)
	if err != nil {
		return w.totalWeight, fmt.Errorf("Trusted add has failed: %w", err)
	}

	w.collectedIDs[signerID] = struct{}{}
	w.totalWeight += info.weight
	return w.totalWeight, nil
}

// TotalWeight returns the total weight presented by the collected signatures.
//
// The function is thread-safe
func (w *WeightedSignatureAggregator) TotalWeight() uint64 {
	w.lock.RLock()
	defer w.lock.RUnlock()
	return w.totalWeight
}

// Aggregate aggregates the signatures and returns the aggregated signature.
//
// The function is thread-safe.
// Aggregate attempts to aggregate the internal signatures and returns the resulting signature data.
// The function performs a final verification and errors if the aggregated signature is not valid. This is
// required for the function safety since "TrustedAdd" allows adding invalid signatures.
//
// TODO : When compacting the list of signers, update the return from []flow.Identifier
// to a compact bit vector.
func (w *WeightedSignatureAggregator) Aggregate() ([]flow.Identifier, []byte, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	// Aggregate includes the safety check of the aggregated signature
	indices, aggSignature, err := w.aggregator.Aggregate()
	if err != nil {
		return nil, nil, fmt.Errorf("aggregate has failed: %w", err)
	}
	signerIDs := make([]flow.Identifier, 0, len(indices))
	for _, index := range indices {
		signerIDs = append(signerIDs, w.ids[index].NodeID)
	}

	return signerIDs, aggSignature, nil
}
