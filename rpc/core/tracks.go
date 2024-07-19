package core

import (
	"encoding/json"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	tracksTypes "github.com/cometbft/cometbft/types/tracks"
	"strconv"
)

func (env *Environment) TracksGetPodCount(_ *rpctypes.Context) (int, error) {
	res, err := env.TxIndexer.GetbytedataFortracks([]byte("countPods"))
	if err != nil {
		return 0, err

	}
	podCount, err := strconv.Atoi(string(res))
	if err != nil {
		return 0, err
	}
	return podCount, nil
}

func (env *Environment) TracksGetPodTxs(_ *rpctypes.Context, podNumber int) ([]tracksTypes.CircuitTransaction, error) {

	env.Logger.Info("Tracks API Request", "req", "tracks_get_pod")

	byteRes, err := env.TxIndexer.GetbytedataFortracks([]byte("raw_pod_" + strconv.Itoa(podNumber)))
	if err != nil {
		return nil, err
	}

	// Deserialize the pod into []tracksTypes.EthTransaction
	var transactionsArrayByte [][]byte
	err = json.Unmarshal(byteRes, &transactionsArrayByte)
	if err != nil {
		return nil, err
	}

	var txArray []tracksTypes.CircuitTransaction
	for _, txByte := range transactionsArrayByte {
		var tx tracksTypes.CircuitTransaction
		err = json.Unmarshal(txByte, &tx)
		if err != nil {
			return nil, err
		}
		txArray = append(txArray, tx)
	}

	return txArray, nil
}
