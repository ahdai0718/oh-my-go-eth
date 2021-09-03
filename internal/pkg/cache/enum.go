package cache

// Type .
type Type int

const (
	_ Type = iota
	// Redis .
	Redis
	// RedisCluster .
	RedisCluster
	// Memcached .
	Memcached
)
