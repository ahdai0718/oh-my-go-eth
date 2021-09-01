package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/glog"
)

func getLatestNBlock(n int) (blockList []*types.Block, err error) {

	blockList = make([]*types.Block, 0)

	client, err := ethclient.Dial(dataSeedURL)

	if err != nil {
		return
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
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
		glog.Infof("====Transaction [%s]====", transaction.Hash().Hex())
	}
}
