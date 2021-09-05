package service

import (
	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/cache"
	"github.com/ahdai0718/oh-my-go-eth/internal/pkg/store"
)

func Init() {
	if err := cache.Init(cache.Redis); err != nil {
		panic(err)
	}

	if err := store.Init(store.StoreTypeMySQLWithCache); err != nil {
		panic(err)
	}
}
