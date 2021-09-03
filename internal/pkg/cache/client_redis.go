package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type clientRedis struct {
	clientBase
	ctx    context.Context
	client *redis.Client
}

// Init .
func (client *clientRedis) Init(serverConfigList []ServerConfig) error {

	if len(serverConfigList) == 0 {
		return fmt.Errorf("sever config is empty")
	}

	client.ctx = context.Background()

	options := &redis.Options{}
	serverConfig := serverConfigList[0]

	options.Addr = fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	options.Password = serverConfig.Password

	client.client = redis.NewClient(options)

	statusCmd := client.client.Ping(client.ctx)

	return statusCmd.Err()
}

// Set .
func (client *clientRedis) Set(key string, value interface{}) error {
	statusCmd := client.client.Set(client.ctx, key, value, 0)
	return statusCmd.Err()
}

// Get .
func (client *clientRedis) Get(key string, value interface{}) error {
	statusCmd := client.client.Get(client.ctx, key)
	err := statusCmd.Scan(value)
	return err
}

func (client *clientRedis) Delete(key string) error {
	statusCmd := client.client.Del(context.Background(), key)
	return statusCmd.Err()
}

// Push .
func (client *clientRedis) Push(key string, values ...interface{}) error {
	return nil
}

// Pop .
func (client *clientRedis) Pop(key string, value interface{}) error {
	return nil
}

// PopAll .
func (client *clientRedis) PopAll(key string, value interface{}) error {
	return nil
}
