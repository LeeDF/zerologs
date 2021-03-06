// Code generated by goctl. DO NOT EDIT!
// Source: helloService.proto

package server

import (
	"context"

	"github.com/LeeDF/zerologs/rpc/helloService/helloService"
	"github.com/LeeDF/zerologs/rpc/helloService/internal/logic"
	"github.com/LeeDF/zerologs/rpc/helloService/internal/svc"
)

type HelloServiceServer struct {
	svcCtx *svc.ServiceContext
}

func NewHelloServiceServer(svcCtx *svc.ServiceContext) *HelloServiceServer {
	return &HelloServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *HelloServiceServer) Ping(ctx context.Context, in *helloService.PingRequest) (*helloService.PingResponse, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
