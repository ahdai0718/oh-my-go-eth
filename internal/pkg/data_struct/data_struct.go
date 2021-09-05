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
		block := NewBlock()
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
		block := NewBlock()
		block.Block = b
		wrapper.List[index] = block
	}
}

type BlockList []*Block

func NewBlock() *Block {
	return &Block{
		Block: &pb.Block{},
	}
}

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

func NewTransaction() *Transaction {
	return &Transaction{
		Transaction: &pb.Transaction{},
	}
}

type Transaction struct {
	*pb.Transaction
}

func (t *Transaction) Parse(transaction *types.Transaction) {
	if transaction == nil {
		glog.Warning("transaction is nil")
		return
	}

	t.TxHash = transaction.Hash().Hex()

	message, err := transaction.AsMessage(types.NewEIP155Signer(transaction.ChainId()), nil)
	if err != nil {
		glog.Error(err)
		return
	}

	t.From = message.From().Hex()
	if message.To() != nil {
		t.To = message.To().Hex()
	}
	t.Nonce = int64(transaction.Nonce())
	t.Value = transaction.Value().String()

}

func NewTransactionLog() *TransactionLog {
	return &TransactionLog{
		TransactionLog: &pb.TransactionLog{},
	}
}

type TransactionLog struct {
	*pb.TransactionLog
}

func (tl *TransactionLog) Parse(log *types.Log) {
	if log == nil {
		glog.Warning("log is nil")
		return
	}

	tl.TxHash = log.TxHash.Hex()
	tl.Index = int64(log.Index)
	tl.Data = string(log.Data)
	tl.TxIndex = int64(log.TxIndex)
}
