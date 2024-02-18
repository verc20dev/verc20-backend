package web3

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fxamacker/cbor/v2"
	"math/big"
)

func GetWeb3Client(rpcUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetLatestBlock(client *ethclient.Client) (uint64, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

func GetTxsInBlock(client *ethclient.Client, blockNumber uint64) ([]types.Transaction, error) {
	block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, err
	}
	txs := make([]types.Transaction, 0)
	for _, tx := range block.Transactions() {
		txs = append(txs, *tx)
	}
	return txs, nil
}

func GetTxFrom(tx *types.Transaction) (string, error) {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	return from.String(), err
}

func SniffIfCbor(tx *types.Transaction) bool {
	if len(tx.Data()) < 3 {
		return false
	}
	cborTag := []byte{0xd9, 0xd9, 0xf7}
	return bytes.Equal(tx.Data()[0:3], cborTag)
}

func CborToJsonString(data []byte) (string, error) {
	var target map[string]interface{}
	err := cbor.Unmarshal(data, &target)
	if err != nil {
		return "", err
	}
	jsonBytes, err := json.Marshal(target)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func GetTxReceipt(client *ethclient.Client, txHash string) (*types.Receipt, error) {
	txHashObj := common.HexToHash(txHash)
	receipt, err := client.TransactionReceipt(context.Background(), txHashObj)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func GetTx(client *ethclient.Client, txHash string) (*types.Transaction, error) {
	txHashObj := common.HexToHash(txHash)
	tx, _, err := client.TransactionByHash(context.Background(), txHashObj)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
