package eth

import (
	"math"
	"sync"
	"time"

	datastruct "github.com/ahdai0718/oh-my-go-eth/internal/app/server/eth/data_struct"
	"github.com/ahdai0718/oh-my-go-eth/internal/app/server/eth/store"
	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/pb"
	"github.com/golang/glog"
)

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
			ethBlock, err := getBlockByNumber(blockNumber)
			if err != nil {
				glog.Error(err)
			}

			block := new(datastruct.Block)
			block.Block = new(pb.Block)
			if ethBlock != nil {
				block.Parse(ethBlock)
				store.DefaultStorer().AddBlock(block.Block)
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
