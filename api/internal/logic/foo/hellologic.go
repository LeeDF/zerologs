package logic

import (
	"context"
	"github.com/LeeDF/zerologs/rpc/helloService/helloserviceclient"

	"github.com/LeeDF/zerologs/api/internal/svc"
	"github.com/LeeDF/zerologs/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) HelloLogic {
	return HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req types.FooReq) (*types.FooRsp, error) {
	ping, err := l.svcCtx.HelloRpc.Ping(l.ctx, &helloserviceclient.PingRequest{Name: req.Name})
	if err != nil {
		return &types.FooRsp{}, err
	}
	return &types.FooRsp{Msg: ping.Msg}, nil
}
