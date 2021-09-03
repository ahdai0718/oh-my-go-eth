package eth

import "github.com/ahdai0718/oh-my-go-eth/internal/app/server/eth/store"

var (
	dataSeedURL string
	scanLimit   = uint(20)
)

func Init() {
	if err := store.DefaultStorer().Init(); err != nil {
		panic(err)
	}
}

func Scan() {
	go scanBlockLoop(0, scanLimit)
}

func SetDataSeedURL(url string) {
	dataSeedURL = url
}

func SetScanLimit(n uint) {
	scanLimit = n
}
