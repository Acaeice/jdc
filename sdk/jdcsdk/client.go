package jdcsdk

import (
	"errors"
	"strings"

	"github.com/gotit/go-net/net"
)

const (
	respOk      = "OK"
	respFail    = "FAIL"
	respUnLogin = "UNLOGIN"
)

var (
	// host 服务地址
	host string
	// Client 服务的客户端
	client *apiClient
)

// API api
type apiClient struct {
	Net    *net.Net
	Header map[string]string
}

func netError(err error) error {
	message := err.Error()
	message = strings.Replace(message, "\"", "", -1)
	return errors.New(message)
}
