package store

import (
	"errors"
	"fmt"
	"time"

	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/pb"
	"github.com/golang/glog"
	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	defaultSimpleFactory = &SimpleFactory{}
)

type Block struct {
	BlockNum   uint64
	BlockHash  string `gorm:"primaryKey"`
	BlockTime  uint64
	ParentHash string
}

func (m Block) TableName() string {
	return "block"
}

func (m *Block) Parse(b *pb.Block) {
	copier.Copy(m, b)
}

func (m *Block) ToPB() (pbBlock *pb.Block) {
	pbBlock = new(pb.Block)
	copier.Copy(pbBlock, m)
	return
}

type Transaction struct {
	TxHash    string `gorm:"primaryKey"`
	BlockHash string
}

func (m Transaction) TableName() string {
	return "transaction"
}

func (m *Transaction) Parse(t *pb.Transaction) {
	copier.Copy(m, t)
}

func (m *Transaction) ToPB() (pbTransaction *pb.Transaction) {
	pbTransaction = new(pb.Transaction)
	copier.Copy(pbTransaction, m)
	return
}

type StorerMySQL struct {
	db *gorm.DB
}

func (storer *StorerMySQL) Init() (err error) {

	host := viper.GetString("database_host")
	port := viper.GetString("database_port")
	username := viper.GetString("database_username")
	passowrd := viper.GetString("database_password")
	schema := viper.GetString("database_schema")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, passowrd, host, port, schema)
	glog.Infof("connecting to database[%s]", dsn)
	storer.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	retry := 0

	if err != nil {
		glog.Errorln(err)

		for err != nil {
			if retry >= 60 {
				return errors.New("over retry to connect to database")
			}
			retry++
			glog.Infof("retry to connect to database[%d]", retry)
			storer.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			time.Sleep(time.Second)
		}
	}

	return
}

func (storer *StorerMySQL) AddBlock(pbBlock *pb.Block) (err error) {
	block := new(Block)
	block.Parse(pbBlock)
	err = storer.db.Create(block).Error
	return
}

func (storer *StorerMySQL) RemoveBlock(pbBlock *pb.Block) (err error) {
	block := new(Block)
	block.Parse(pbBlock)
	err = storer.db.Delete(block).Error
	return
}

func (storer *StorerMySQL) GetBlockByNum(blockNumber uint64) (pbBlock *pb.Block, err error) {

	block := new(Block)
	block.BlockNum = blockNumber

	err = storer.db.Take(block).Error
	if err != nil {
		return
	}

	pbBlock = block.ToPB()

	return
}

func (storer *StorerMySQL) GetBlockByHash(blockHash string) (pbBlock *pb.Block, err error) {
	block := new(Block)
	block.BlockHash = blockHash

	err = storer.db.Take(block).Error
	if err != nil {
		return
	}

	pbBlock = block.ToPB()
	return
}

func (storer *StorerMySQL) GetLatestNBlock(n uint) (pbBlockList []*pb.Block, err error) {
	pbBlockList = make([]*pb.Block, n)
	blockList := make([]*Block, n)

	err = storer.db.Order("block_num DESC").Limit(int(n)).Find(blockList).Error
	if err != nil {
		return
	}

	for index, block := range blockList {
		pbBlockList[index] = block.ToPB()
	}

	return
}

func (storer *StorerMySQL) AddTransaction(pbTransaction *pb.Transaction) (err error) {
	transaction := new(Transaction)
	transaction.Parse(pbTransaction)
	err = storer.db.Create(transaction).Error
	return
}

func (storer *StorerMySQL) GetTransactionByHash(txHash string) (pbTransaction *pb.Transaction, err error) {
	transaction := new(Transaction)
	transaction.TxHash = txHash

	err = storer.db.Take(transaction).Error
	if err != nil {
		return
	}

	pbTransaction = transaction.ToPB()
	return
}

func (storer *StorerMySQL) GetTransactionListByBlockHash(blockHash string) (pbTransactionList []*pb.Transaction, err error) {

	transctionList := make([]*Transaction, 0)

	err = storer.db.Where("block_hash = ?", blockHash).Find(&transctionList).Error
	if err != nil {
		return
	}

	if len(transctionList) > 0 {
		pbTransactionList = make([]*pb.Transaction, len(transctionList))
		for index, transction := range transctionList {
			pbTransaction := transction.ToPB()
			pbTransactionList[index] = pbTransaction
		}
	}

	return
}
