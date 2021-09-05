package cache

import "time"

type clientBase struct {
}

// Init .
func (client *clientBase) Init() error {
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
