//go:build gofuzz || go1.21

package tests

import (
	"testing"

	abciclient "github.com/airchains-network/wasmbft/abci/client"
	"github.com/airchains-network/wasmbft/abci/example/kvstore"
	"github.com/airchains-network/wasmbft/config"
	cmtsync "github.com/airchains-network/wasmbft/libs/sync"
	mempool "github.com/airchains-network/wasmbft/mempool"
)

func FuzzMempool(f *testing.F) {
	app := kvstore.NewInMemoryApplication()
	mtx := new(cmtsync.Mutex)
	conn := abciclient.NewLocalClient(mtx, app)
	err := conn.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mp := mempool.NewCListMempool(cfg, conn, 0)

	f.Fuzz(func(t *testing.T, data []byte) {
		_ = mp.CheckTx(data, nil, mempool.TxInfo{})
	})
}
