package cache

import "github.com/golang/glog"

// Create .
func Create(t Type) (client Client) {
	switch t {
	case Redis:
		client = &clientRedis{}
	}

	if client == nil {
		glog.Errorf("Type [%d] is not implement", t)
	}

	return
}
