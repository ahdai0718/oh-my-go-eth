package cache

import (
	"time"
)

const (
	ERROR_REDIS_NO_DATA = "redis: nil"
	RetryLimit          = 60
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
	Init() error
	Set(key string, value interface{}) error
	SetWithExpiration(expiration time.Duration, key string, value interface{}) error
	Get(key string, value interface{}) error
	Delete(key string) error
}
