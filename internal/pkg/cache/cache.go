package cache

import (
	"time"
)

const (
	ERROR_REDIS_NO_DATA = "redis: nil"
)

// ServerConfig ...
type ServerConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

// Client ...
type Client interface {
	Init(serverConfigList []ServerConfig) error
	Set(key string, value interface{}) error
	SetWithExpiration(expiration time.Duration, key string, value interface{}) error
	Get(key string, value interface{}) error
	Delete(key string) error

	Push(key string, values ...interface{}) error
	PushWithExpiration(expiration time.Duration, key string, values ...interface{}) error
	Pop(key string, value interface{}) error
	PopAll(key string, value interface{}) error

	HashSet(key string, field string, value interface{}) error
	HashGet(key string, field string, value interface{}) error
	HashGetAll(key string, value interface{}) error
	HashDelete(key string, field string) error
}
