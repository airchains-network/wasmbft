package null

import (
	"context"
	"errors"

	"github.com/airchains-network/wasmbft/libs/log"

	abci "github.com/airchains-network/wasmbft/abci/types"
	"github.com/airchains-network/wasmbft/libs/pubsub/query"
	"github.com/airchains-network/wasmbft/state/txindex"
)

var _ txindex.TxIndexer = (*TxIndex)(nil)

// TxIndex acts as a /dev/null.
type TxIndex struct{}

// Get on a TxIndex is disabled and panics when invoked.
func (txi *TxIndex) Get(_ []byte) (*abci.TxResult, error) {
	return nil, errors.New(`indexing is disabled (set 'tx_index = "kv"' in config)`)
}

// AddBatch is a noop and always returns nil.
func (txi *TxIndex) AddBatch(_ *txindex.Batch) error {
	return nil
}

// Index is a noop and always returns nil.
func (txi *TxIndex) Index(_ *abci.TxResult) error {
	return nil
}

func (txi *TxIndex) Search(_ context.Context, _ *query.Query) ([]*abci.TxResult, error) {
	return []*abci.TxResult{}, nil
}

func (txi *TxIndex) SetLogger(log.Logger) {

}
