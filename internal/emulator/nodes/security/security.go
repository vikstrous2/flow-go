package security

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/dapperlabs/bamboo-node/internal/emulator/data"
	"github.com/dapperlabs/bamboo-node/internal/emulator/nodes/security/block_builder"
)

// Node simulates the behaviour of a Bamboo security node.
type Node struct {
	conf          *Config
	state         *data.WorldState
	collectionsIn chan *data.Collection
	blockBuilder  *block_builder.BlockBuilder
	log           *logrus.Logger
}

// Config hold the configuration options for an security node.
type Config struct {
	BlockInterval time.Duration
}

// NewNode returns a new simulated security node.
func NewNode(conf *Config, state *data.WorldState, collectionsIn chan *data.Collection, log *logrus.Logger) *Node {
	blockBuilder := block_builder.NewBlockBuilder(state, collectionsIn, log)

	return &Node{
		conf:          conf,
		state:         state,
		collectionsIn: collectionsIn,
		blockBuilder:  blockBuilder,
		log:           log,
	}
}

func (n *Node) Start(ctx context.Context) {
	n.blockBuilder.Start(ctx, n.conf.BlockInterval)
}
