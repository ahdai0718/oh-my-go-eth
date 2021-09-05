package cache

import (
	"github.com/golang/glog"
)

var (
	client Client
)

// Init .
func Init(t Type) (err error) {

	client = Create(t)

	err = client.Init()
	if err != nil {
		glog.Error(err)
		return err
	}

	err = client.Set("test", "test")
	if err != nil {
		glog.Error(err)
		return err
	}

	return
}

// DefaultClient .
func DefaultClient() Client {
	return client
}
