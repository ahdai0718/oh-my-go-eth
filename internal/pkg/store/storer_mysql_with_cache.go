package store

import (
	"fmt"

	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/cache"
	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type StorerMySQLWithCache struct {
	*StorerMySQL
}

func (storer *StorerMySQLWithCache) getFromCache(key string, m protoreflect.ProtoMessage) (err error) {
	var data string
	err = cache.DefaultClient().Get(key, &data)
	if err == nil && data != "" {
		err = proto.Unmarshal([]byte(data), m)
		if err == nil {
			return
		}
	}
	return
}

func (storer *StorerMySQLWithCache) setToCache(key string, m protoreflect.ProtoMessage) (err error) {
	var data []byte
	data, err = proto.Marshal(m)
	if err == nil && len(data) > 0 {
		err = cache.DefaultClient().Set(key, string(data))
		if err == nil {
			return
		}
	}
	return
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
	pbBlock = new(pb.Block)

	cacheKey := fmt.Sprintf("BlockByNum%d", blockNumber)
	err = storer.getFromCache(cacheKey, pbBlock)
	if err == nil {
		return
	}

	if cache.DefaultClient().IsNotFound(err) {
		pbBlock, err = storer.StorerMySQL.GetBlockByNum(blockNumber)
		if err == nil && pbBlock != nil {
			err = storer.setToCache(cacheKey, pbBlock)
			if err == nil {
				return
			}
		}
	}

	return storer.StorerMySQL.GetBlockByNum(blockNumber)
}

func (storer *StorerMySQLWithCache) GetBlockByHash(blockHash string) (pbBlock *pb.Block, err error) {

	pbBlock = new(pb.Block)

	cacheKey := fmt.Sprintf("BlockByHash%s", blockHash)
	err = storer.getFromCache(cacheKey, pbBlock)
	if err == nil {
		return
	}

	if cache.DefaultClient().IsNotFound(err) {
		pbBlock, err = storer.StorerMySQL.GetBlockByHash(blockHash)
		if err == nil && pbBlock != nil {
			err = storer.setToCache(cacheKey, pbBlock)
			if err == nil {
				return
			}
		}
	}

	return storer.StorerMySQL.GetBlockByHash(blockHash)
}

func (storer *StorerMySQLWithCache) GetLatestNBlock(n uint) (pbBlockList []*pb.Block, err error) {

	lastedtBlockList, err := storer.StorerMySQL.GetLatestNBlock(1)
	if err != nil {
		return
	}

	lastedBlock := lastedtBlockList[0]

	pbBlockList = make([]*pb.Block, 0)
	pbBlockList = append(pbBlockList, lastedBlock)

	parentBlockHash := lastedBlock.ParentHash

	for len(pbBlockList) < int(n) {
		var pbBlock *pb.Block
		pbBlock, err = storer.GetBlockByHash(parentBlockHash)
		if err != nil {
			return
		}
		pbBlockList = append(pbBlockList, pbBlock)
		parentBlockHash = pbBlock.ParentHash
	}

	return

}

func (storer *StorerMySQLWithCache) AddTransaction(pbTransaction *pb.Transaction) (err error) {

	return storer.StorerMySQL.AddTransaction(pbTransaction)
}

func (storer *StorerMySQLWithCache) GetTransactionByHash(txHash string) (pbTransaction *pb.Transaction, err error) {

	pbTransaction = new(pb.Transaction)

	cacheKey := fmt.Sprintf("TransactionByHash%s", txHash)
	err = storer.getFromCache(cacheKey, pbTransaction)
	if err == nil {
		return
	}

	if cache.DefaultClient().IsNotFound(err) {
		pbTransaction, err = storer.StorerMySQL.GetTransactionByHash(txHash)
		if err == nil && pbTransaction != nil {
			err = storer.setToCache(cacheKey, pbTransaction)
			if err == nil {
				return
			}
		}
	}

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
