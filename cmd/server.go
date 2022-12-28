package main

import (
	"flag"
	"github.com/mix-plus/grpc-skeleton/internal/server"

	"github.com/go-ll/mrpc"
	"github.com/mix-plus/core/conf"
	"github.com/mix-plus/grpc-skeleton/api"
	"github.com/mix-plus/grpc-skeleton/internal/config"
	"github.com/mix-plus/grpc-skeleton/internal/svc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c)
	// newGrpcServer
	srv := server.NewHelloServer(ctx)
	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		// grpc register
		hello.RegisterHelloServer(g, srv)
	})

	defer s.Stop()

	s.Start()
}
