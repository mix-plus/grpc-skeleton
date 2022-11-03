package main

import (
	"flag"
	"github.com/go-ll/mrpc"
	"github.com/mix-plus/core/conf"
	"github.com/mix-plus/grpc-skeleton/internal/config"
	"github.com/mix-plus/grpc-skeleton/internal/svc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/demo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	_ = svc.NewServiceContext(c)
	// newGrpcServer

	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		// grpc register

	})

	defer s.Stop()

	s.Start()
}
