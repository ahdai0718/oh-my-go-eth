package store

import (
	"fmt"

	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/cache"
	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/pb"
	"google.golang.org/protobuf/proto"
)

type StorerMySQLWithCache struct {
	*StorerMySQL
}

func (storer *StorerMySQLWithCache) Init() (err error) {

	return storer.StorerMySQL.Init()
}

func (storer *StorerMySQLWithCache) AddBlock(pbBlock *pb.Block) (err error) {

	return storer.StorerMySQL.AddBlock(pbBlock)
}

func (storer *StorerMySQLWithCache) RemoveBlock(pbBlock *pb.Block) (err error) {

	return storer.StorerMySQL.RemoveBlock(pbBlock)
}

func (storer *StorerMySQLWithCache) GetBlockByNum(blockNumber uint64) (pbBlock *pb.Block, err error) {

	return storer.StorerMySQL.GetBlockByNum(blockNumber)
}

func (storer *StorerMySQLWithCache) GetBlockByHash(blockHash string) (pbBlock *pb.Block, err error) {

	cacheKey := fmt.Sprintf("BlockByHash%s", blockHash)
	data := ""
	err = cache.DefaultClient().Get(cacheKey, &data)
	if err == nil && data != "" {
		err = proto.Unmarshal([]byte(data), pbBlock)
		if err == nil {
			return
		}
	}

	return storer.StorerMySQL.GetBlockByHash(blockHash)
}

func (storer *StorerMySQLWithCache) GetLatestNBlock(n uint) (pbBlockList []*pb.Block, err error) {

	return storer.StorerMySQL.GetLatestNBlock(n)
}

func (storer *StorerMySQLWithCache) AddTransaction(pbTransaction *pb.Transaction) (err error) {

	return storer.StorerMySQL.AddTransaction(pbTransaction)
}

func (storer *StorerMySQLWithCache) GetTransactionByHash(txHash string) (pbTransaction *pb.Transaction, err error) {

	return storer.StorerMySQL.GetTransactionByHash(txHash)
}

func (storer *StorerMySQLWithCache) GetTransactionListByBlockHash(blockHash string) (pbTransactionList []*pb.Transaction, err error) {

	return storer.StorerMySQL.GetTransactionListByBlockHash(blockHash)
}

func (storer *StorerMySQLWithCache) AddTransactionLog(pbTransactionLog *pb.TransactionLog) (err error) {

	return storer.StorerMySQL.AddTransactionLog(pbTransactionLog)
}

func (storer *StorerMySQLWithCache) GetTransactionLogListByBlockTxHash(txHash string) (pbTransactionLogList []*pb.TransactionLog, err error) {

	return storer.StorerMySQL.GetTransactionLogListByBlockTxHash(txHash)
}
