package eth

import "github.com/ahdai0718/oh-my-go-eth/internal/app/server/eth/store"

var (
	dataSeedURL string
)

func Init() {
	if err := store.DefaultStorer().Init(); err != nil {
		panic(err)
	}
}

func Run() {

}

func SetDataSeedURL(url string) {
	dataSeedURL = url
}
