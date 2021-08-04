package jdcsdk

import (
	"github.com/gotit/go-net/net"
)

// Init 初始化服务
func Init(serviceHost string) {
	host = serviceHost
	client = &apiClient{
		Net: net.New(),
	}
}
