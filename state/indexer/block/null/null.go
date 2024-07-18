package null

import (
	"context"
	"errors"

	"github.com/airchains-network/wasmbft/libs/log"
	"github.com/airchains-network/wasmbft/libs/pubsub/query"
	"github.com/airchains-network/wasmbft/state/indexer"
	"github.com/airchains-network/wasmbft/types"
)

var _ indexer.BlockIndexer = (*BlockerIndexer)(nil)

// TxIndex implements a no-op block indexer.
type BlockerIndexer struct{}

func (idx *BlockerIndexer) Has(int64) (bool, error) {
	return false, errors.New(`indexing is disabled (set 'tx_index = "kv"' in config)`)
}

func (idx *BlockerIndexer) Index(types.EventDataNewBlockEvents) error {
	return nil
}

func (idx *BlockerIndexer) Search(context.Context, *query.Query) ([]int64, error) {
	return []int64{}, nil
}

func (idx *BlockerIndexer) SetLogger(log.Logger) {
}
