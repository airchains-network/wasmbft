package kv

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/state/txindex"
	"github.com/cometbft/cometbft/types"
	tracksTypes "github.com/cometbft/cometbft/types/tracks"
	"strconv"
	"strings"
)

func extractAttribute(attributes []abci.EventAttribute, key string) string {
	for _, attr := range attributes {
		if string(attr.Key) == key {
			return string(attr.Value)
		}
	}
	return ""
}

func StorePod(txi *TxIndex, b *txindex.Batch) error {

	storeBatch := txi.store.NewBatch()
	defer storeBatch.Close()
	if len(b.Ops) == 0 {
		txi.log.Info("no operations provided", "module", "txindex")
		return nil
	}

	var txArray []tracksTypes.CircuitTransaction
	for _, result := range b.Ops {
		var (
			sender    = ""
			recipient = ""
			amount    = ""
			//gasWanted int64  = 0
			gasUsed     int64 = 0
			txHash            = ""
			fromBalance       = ""
			toBalance         = ""
		)
		events := result.Result.Events
		txHash = hex.EncodeToString(types.Tx(result.Tx).Hash())
		//gasWanted = result.Result.GasWanted
		gasUsed = result.Result.GasUsed
		// events under this transaction.
		for _, event := range events {
			if event.Type == "transfer" {
				attr := event.Attributes
				for _, attribute := range attr {
					if string(attribute.Key) == "recipient" {
						recipient = string(attribute.Value)
					}
					if string(attribute.Key) == "sender" {
						sender = string(attribute.Value)
					}
					if string(attribute.Key) == "amount" {
						amount = string(attribute.Value)
					}
					if string(attribute.Key) == "from_balance" {
						fromBalance = string(attribute.Value)
					}
					if string(attribute.Key) == "to_balance" {
						toBalance = string(attribute.Value)
					}
				}

				// Retrieve the current nonce for the sender
				nonce, err := GetNonce(txi, sender)
				if err != nil {
					return fmt.Errorf("failed to retrieve nonce: %w", err)
				}
				nonce++
				err = SetNonce(txi, sender, nonce)
				if err != nil {
					return fmt.Errorf("failed to store nonce: %w", err)
				}

				amounts := strings.Split(amount, ",")
				fromBalances := strings.Split(fromBalance, ",")
				toBalances := strings.Split(toBalance, ",")
				//recipients := strings.Split(recipient, ",")

				for i := 0; i < len(amounts); i++ {
					txArray = append(txArray, tracksTypes.CircuitTransaction{
						Nonce:       strconv.FormatUint(nonce-1, 10),
						From:        sender,
						To:          recipient,
						Amount:      amounts[i],
						TxHash:      txHash,
						Gas:         strconv.FormatInt(gasUsed, 10),
						FromBalance: fromBalances[i],
						ToBalance:   toBalances[i],
					})
				}
			}
		}
	}

	// Retrieve the current pod number
	podNumber, err := RetrievePodCount(txi)
	if err != nil {
		return fmt.Errorf("failed to retrieve pod count: %w", err)
	}

	// Retrieve the current pod if it exists
	var currentPod [][]byte
	currentPod, err = GetPod(txi, podNumber)
	if err != nil && err.Error() != "pod not found" {
		return fmt.Errorf("failed to retrieve current pod: %w", err)
	}

	txi.log.Info("processing pod", "podNumber", podNumber, "txLen", len(currentPod), "module", "txindex")

	// Process transactions in txArray for pods storage
	for len(txArray) > 0 {
		if len(currentPod) < transactionPodSize {
			// Fill the current pod
			for len(currentPod) < transactionPodSize && len(txArray) > 0 {
				txBytes, err := json.Marshal(txArray[0])
				if err != nil {
					return fmt.Errorf("failed to serialize transaction: %w", err)
				}
				currentPod = append(currentPod, txBytes)
				txArray = txArray[1:]
			}

			// If the current pod is filled, store it
			if len(currentPod) == transactionPodSize {
				err = SetPod(txi, podNumber, currentPod)
				if err != nil {
					return fmt.Errorf("failed to store pod: %w", err)
				}

				// Increment pod count
				err = IncrementPodCount(txi)
				if err != nil {
					return fmt.Errorf("failed to increment pod count: %w", err)
				}

				// Reset current pod and increment pod number
				currentPod = [][]byte{}
				podNumber++
			}
		}
	}

	// Store remaining transactions in the current pod if not empty
	if len(currentPod) > 0 {
		err = SetPod(txi, podNumber, currentPod)
		if err != nil {
			return fmt.Errorf("failed to store incomplete pod: %w", err)
		}
	}

	return nil
}
