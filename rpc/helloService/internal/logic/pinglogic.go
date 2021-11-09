package logic

import (
	"context"

	"github.com/LeeDF/zerologs/rpc/helloService/helloService"
	"github.com/LeeDF/zerologs/rpc/helloService/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *helloService.PingRequest) (*helloService.PingResponse, error) {
	//var m map[string]string
	//m[in.Name] = in.Name
	panic("test")
	return &helloService.PingResponse{Msg: "hello " + in.Name}, nil
}
