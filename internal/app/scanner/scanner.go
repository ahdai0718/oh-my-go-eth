package scanner

import (
	"math"
	"sync"
	"time"

	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/cache"
	datastruct "github.com/ahdai0718/oh-my-go-eth/internal/pkg/data_struct"
	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/eth"
	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/pb"
	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/store"
	"github.com/golang/glog"
)

var (
	scanLimit = uint(20)
)

func Init(dataSeedURL string) {
	if err := cache.Init(cache.Redis); err != nil {
		panic(err)
	}

	if err := store.Init(store.StoreTypeMySQLWithCache); err != nil {
		panic(err)
	}

	if err := eth.Init(dataSeedURL); err != nil {
		panic(err)
	}

}

func Scan() {
	go scanBlockLoop(0, scanLimit)
}

func SetScanLimit(n uint) {
	scanLimit = n
}

func scanBlockLoop(startNumber uint64, limit uint) {
	nextNumber := startNumber
	for {
		nextNumber = scanNBlock(nextNumber, limit)
		time.Sleep(time.Second)
	}
}

func scanNBlock(startNumber uint64, limit uint) (nextNumber uint64) {
	wg := sync.WaitGroup{}

	maxNumber := startNumber + uint64(limit)

	blockChan := make(chan *datastruct.Block, limit)

	wg.Add(int(limit))

	for number := startNumber; number < maxNumber; number++ {

		go func(blockNumber uint64) {
			defer wg.Done()
			ethBlock, err := eth.GetBlockByNumber(blockNumber)
			if err != nil {
				glog.Error(err)
			}

			block := datastruct.NewBlock()
			block.Block = new(pb.Block)
			if ethBlock != nil {
				block.Parse(ethBlock)
				err := store.DefaultStorer().AddBlock(block.Block)
				if err != nil {
					glog.Error(err)
				}

				transactions := ethBlock.Transactions()
				if len(transactions) > 0 {
					for _, ethTransaction := range transactions {

						ethTransactionDetail, err := eth.GetTransactionByHash(ethTransaction.Hash().Hex())
						if err != nil {
							glog.Error(err)
						}

						if ethTransactionDetail != nil {
							transaction := datastruct.NewTransaction()
							transaction.BlockHash = block.BlockHash
							transaction.Parse(ethTransactionDetail)

							err := store.DefaultStorer().AddTransaction(transaction.Transaction)
							if err != nil {
								glog.Error(err)
							}

							ethTransactionReceipt, err := eth.GetTransactionReceiptByHash(ethTransaction.Hash().Hex())
							if err != nil {
								glog.Error(err)
							}

							if ethTransactionReceipt != nil {
								for _, ethTransactionLog := range ethTransactionReceipt.Logs {
									transactionLog := datastruct.NewTransactionLog()
									transactionLog.TxHash = transaction.TxHash
									transactionLog.Parse(ethTransactionLog)

									err := store.DefaultStorer().AddTransactionLog(transactionLog.TransactionLog)
									if err != nil {
										glog.Error(err)
									}
								}
							}
						}
					}
				}
			}

			blockChan <- block
		}(number)
	}

	wg.Wait()
	close(blockChan)

	nextNumber = 0
	for block := range blockChan {
		nextNumber = uint64(math.Max(float64(nextNumber), float64(block.BlockNum)))
	}
	nextNumber = nextNumber + 1

	return
}
