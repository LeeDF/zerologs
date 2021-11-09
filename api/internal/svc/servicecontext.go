package svc

import (
	"github.com/LeeDF/zerologs/api/internal/config"
	"github.com/LeeDF/zerologs/rpc/helloService/helloserviceclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	HelloRpc helloserviceclient.HelloService
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, _ := zrpc.NewClient(zrpc.NewDirectClientConf([]string{"127.0.0.1:8080"}, "", ""))
	return &ServiceContext{
		Config:   c,
		HelloRpc: helloserviceclient.NewHelloService(client),
	}
}
