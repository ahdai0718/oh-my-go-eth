package store

import "github.com/ahdai0718/oh-my-go-eth/internal/pkg/pb"

var (
	defaultStorer = defaultSimpleFactory.Create(StoreTypeMySQL)
)

type Storer interface {
	Init() (err error)
	AddBlock(pbBlock *pb.Block) (err error)
	RemoveBlock(pbBlock *pb.Block) (err error)
	GetBlockByNum(blockNumber uint64) (pbBlock *pb.Block, err error)
	GetBlockByHash(blockHash string) (pbBlock *pb.Block, err error)
	GetLatestNBlock(n uint) (pbBlockList []*pb.Block, err error)
	AddTransaction(pbTransaction *pb.Transaction) (err error)
	GetTransactionByHash(txHash string) (pbTransaction *pb.Transaction, err error)
	GetTransactionListByBlockHash(blockHash string) (pbTransactionList []*pb.Transaction, err error)
}

func DefaultStorer() Storer {
	return defaultStorer
}
