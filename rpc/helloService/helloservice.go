package main

import (
	"flag"
	"fmt"

	"github.com/LeeDF/zerologs/rpc/helloService/helloService"
	"github.com/LeeDF/zerologs/rpc/helloService/internal/config"
	"github.com/LeeDF/zerologs/rpc/helloService/internal/server"
	"github.com/LeeDF/zerologs/rpc/helloService/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/helloservice.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)
	ctx := svc.NewServiceContext(c)
	srv := server.NewHelloServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		helloService.RegisterHelloServiceServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
