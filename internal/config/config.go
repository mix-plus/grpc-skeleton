package config

import (
	"github.com/go-ll/mrpc"
)

type Config struct {
	mrpc.RpcServerConf `mapstructure:",squash"`
}
