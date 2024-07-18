package core

import (
	ctypes "github.com/airchains-network/wasmbft/rpc/core/types"
	rpctypes "github.com/airchains-network/wasmbft/rpc/jsonrpc/types"
)

// Health gets node health. Returns empty result (200 OK) on success, no
// response - in case of an error.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/Info/health
func (env *Environment) Health(*rpctypes.Context) (*ctypes.ResultHealth, error) {
	return &ctypes.ResultHealth{}, nil
}
