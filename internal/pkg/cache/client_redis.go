package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

type clientRedis struct {
	clientBase
	ctx    context.Context
	client *redis.Client
}

// Init .
func (client *clientRedis) Init() (err error) {

	serverConfig := &ServerConfig{
		Host: viper.GetString("redis_host"),
		Port: viper.GetInt("redis_port"),
	}

	client.ctx = context.Background()

	options := &redis.Options{}

	options.Addr = fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	options.Password = serverConfig.Password

	client.client = redis.NewClient(options)
	statusCmd := client.client.Ping(client.ctx)
	err = statusCmd.Err()

	retry := 0

	if err != nil {
		glog.Errorln(err)

		for err != nil {
			if retry >= RetryLimit {
				return errors.New("over retry to connect to redis")
			}
			retry++
			glog.Infof("retry to connect to redis[%d]", retry)
			statusCmd = client.client.Ping(client.ctx)
			err = statusCmd.Err()
			if err != nil {
				glog.Errorln(err)
				time.Sleep(time.Second)
			}
		}
	}

	return
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
