package server

import (
	"context"

	hello "github.com/mix-plus/grpc-skeleton/api"
	"github.com/mix-plus/grpc-skeleton/internal/svc"
)

type HelloServer struct {
	svcCtx *svc.ServiceContext
}

func NewHelloServer(ctx *svc.ServiceContext) *HelloServer {
	return &HelloServer{
		svcCtx: ctx,
	}
}

func (h HelloServer) SayHello(ctx context.Context, req *hello.HelloReq) (*hello.HelloResp, error) {
	return &hello.HelloResp{}, nil
}
