package cache

import (
	"strconv"

	"github.com/golang/glog"
)

var (
	// Cache
	serverHostList = []string{}
	serverPortList = []string{}
	client         Client
)

// Init .
func Init(t Type) (err error) {

	loadCacheConfig()

	client = Create(t)

	serverConfigList := []ServerConfig{}

	for index, serverHost := range serverHostList {

		strPort := serverPortList[index]

		port, err := strconv.Atoi(strPort)

		if err != nil {
			return err
		}

		serverConfig := ServerConfig{
			Host: serverHost,
			Port: port,
		}
		serverConfigList = append(serverConfigList, serverConfig)
	}

	glog.Info(serverConfigList)

	err = client.Init(serverConfigList)
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
	if client == nil {
		panic("should init cache client first")
	}

	return client
}

func loadCacheConfig() {

}
