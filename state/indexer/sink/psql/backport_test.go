package psql

import (
	"github.com/airchains-network/wasmbft/state/indexer"
	"github.com/airchains-network/wasmbft/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
