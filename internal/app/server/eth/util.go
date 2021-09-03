package eth

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/glog"
)

func getLatestNBlock(n int) (blockList []*types.Block, err error) {

	blockList = make([]*types.Block, 0)

	client, err := ethclient.Dial(dataSeedURL)

	if err != nil {
		glog.Error(err)
		return
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		glog.Error(err)
		return
	}

	printBlockInfo(block)
	blockList = append(blockList, block)

	if n > 0 {
		for len(blockList) < n {
			parentHash := block.ParentHash()

			block, err = client.BlockByHash(context.Background(), parentHash)
			if err != nil {
				return
			}

			printBlockInfo(block)
			blockList = append(blockList, block)
		}
	}

	return
}

func getBlockByNumber(number uint64) (block *types.Block, err error) {

	client, err := ethclient.Dial(dataSeedURL)

	if err != nil {
		return
	}

	block, err = client.BlockByNumber(context.Background(), big.NewInt(int64(number)))
	if err != nil {
		return
	}

	printBlockInfo(block)

	return
}

func getTransactionByHash(txHash string) (err error) {

	client, err := ethclient.Dial(dataSeedURL)

	if err != nil {
		return
	}

	hash := common.Hash{}
	err = hash.UnmarshalText([]byte(txHash))
	if err != nil {
		return
	}

	transaction, _, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return
	}

	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return
	}

	printTransactionInfo(transaction)
	printReceiptInfo(receipt)

	return
}

func printBlockInfo(block *types.Block) {
	if block == nil {
		glog.Warning("block is nil")
		return
	}
	glog.Infoln("====Block====")
	glog.Infof("Number:%d", block.Number().Uint64())
	glog.Infof("GasLimit:%d", block.GasLimit())
	glog.Infof("GasUsed:%d", block.GasUsed())
	glog.Infof("Difficulty:%d", block.Difficulty().Uint64())
	glog.Infof("Time:%d", block.Time())
	glog.Infof("Hash:%s", block.Hash().Hex())
	glog.Infof("ParentHash:%s", block.ParentHash().Hex())
	glog.Infof("UncleHash:%s", block.UncleHash().Hex())

	glog.Infof("====Transactions[%d]====", len(block.Transactions()))
	for _, transaction := range block.Transactions() {
		printTransactionInfo(transaction)
	}
}

func printTransactionInfo(transaction *types.Transaction) {
	if transaction == nil {
		glog.Warning("transaction is nil")
		return
	}

	message, err := transaction.AsMessage(types.NewEIP155Signer(transaction.ChainId()), nil)
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Infoln("====Transaction====")
	glog.Infof("Hash:%s", transaction.Hash().Hex())
	glog.Infof("From:%s", message.From().Hex())
	if message.To() != nil {
		glog.Infof("To:%s", message.To().Hex())
	}
	glog.Infof("Nonce:%d", transaction.Nonce())
	glog.Infof("Value:%d", transaction.Value().Uint64())
}

func printReceiptInfo(receipt *types.Receipt) {
	if receipt == nil {
		glog.Warning("receipt is nil")
		return
	}

	glog.Infof("==== Receipt ====")

	for _, log := range receipt.Logs {
		printLogInfo(log)
	}
}

func printLogInfo(log *types.Log) {
	if log == nil {
		glog.Warning("log is nil")
		return
	}

	glog.Infof("==== Log ====")
	glog.Infof("Index:%d", log.Index)
	glog.Infof("Data:%s", hex.EncodeToString(log.Data))
	glog.Infof("TxHash:%s", log.TxHash.Hex())
	glog.Infof("TxIndex:%d", log.TxIndex)

}
