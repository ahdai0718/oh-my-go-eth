package datastruct

import (
	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/pb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/glog"
)

type BlockListWrapper struct {
	List BlockList `json:"blocks"`
}

func (wrapper *BlockListWrapper) Parse(blockList []*types.Block) {
	if len(blockList) < 1 {
		return
	}

	wrapper.List = make(BlockList, len(blockList))

	for index, b := range blockList {
		block := new(Block)
		block.Parse(b)
		wrapper.List[index] = block
	}
}

func (wrapper *BlockListWrapper) ParsePB(blockList []*pb.Block) {
	if len(blockList) < 1 {
		return
	}

	wrapper.List = make(BlockList, len(blockList))

	for index, b := range blockList {
		block := new(Block)
		block.Block = b
		wrapper.List[index] = block
	}
}

type BlockList []*Block

type Block struct {
	*pb.Block
}

func (b *Block) Parse(block *types.Block) {
	if block == nil {
		glog.Warning("block is nil")
		return
	}
	b.BlockNum = block.Number().Int64()
	b.BlockHash = block.Hash().Hex()
	b.BlockTime = int64(block.Time())
	b.ParentHash = block.ParentHash().Hex()

	if len(block.Transactions()) > 0 {
		b.Transactions = make([]*pb.Transaction, len(block.Transactions()))
		for index, transactions := range block.Transactions() {
			pbTransaction := new(pb.Transaction)
			pbTransaction.TxHash = transactions.Hash().Hex()
			b.Transactions[index] = pbTransaction
		}
	}
}
