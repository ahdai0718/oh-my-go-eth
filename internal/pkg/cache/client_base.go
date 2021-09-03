package cache

import "time"

type clientBase struct {
}

// Init .
func (client *clientBase) Init(serverConfigList []ServerConfig) error {
	return nil
}

// Set .
func (client *clientBase) Set(key string, value interface{}) error {
	return nil
}

// SetWithExpiration .
func (client *clientBase) SetWithExpiration(expiration time.Duration, key string, value interface{}) error {
	return nil
}

// Get .
func (client *clientBase) Get(key string, value interface{}) error {
	return nil
}

// Get .
func (client *clientBase) Delete(key string) error {
	return nil
}

// Push .
func (client *clientBase) Push(key string, values ...interface{}) error {
	return nil
}

// PushWithExpiration .
func (client *clientBase) PushWithExpiration(expiration time.Duration, key string, values ...interface{}) error {
	return nil
}

// Pop .
func (client *clientBase) Pop(key string, value interface{}) error {
	return nil
}

// PopAll .
func (client *clientBase) PopAll(key string, value interface{}) error {
	return nil
}

// HashSet .
func (client *clientBase) HashSet(key string, field string, value interface{}) error {
	return nil
}

// HashGet .
func (client *clientBase) HashGet(key string, field string, value interface{}) error {
	return nil
}

// HashGetAll .
func (client *clientBase) HashGetAll(key string, value interface{}) (err error) {
	return nil
}

// HashDelete .
func (client *clientBase) HashDelete(key string, field string) error {
	return nil
}
