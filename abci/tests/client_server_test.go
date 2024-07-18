package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	abciclient "github.com/airchains-network/wasmbft/abci/client"
	"github.com/airchains-network/wasmbft/abci/example/kvstore"
	abciserver "github.com/airchains-network/wasmbft/abci/server"
)

func TestClientServerNoAddrPrefix(t *testing.T) {
	t.Helper()

	addr := "localhost:26658"
	transport := "socket"
	app := kvstore.NewInMemoryApplication()

	server, err := abciserver.NewServer(addr, transport, app)
	assert.NoError(t, err, "expected no error on NewServer")
	err = server.Start()
	assert.NoError(t, err, "expected no error on server.Start")
	t.Cleanup(func() {
		if err := server.Stop(); err != nil {
			t.Error(err)
		}
	})

	client, err := abciclient.NewClient(addr, transport, true)
	assert.NoError(t, err, "expected no error on NewClient")
	err = client.Start()
	assert.NoError(t, err, "expected no error on client.Start")
	t.Cleanup(func() {
		if err := client.Stop(); err != nil {
			t.Error(err)
		}
	})
}
