package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	User  zrpc.RpcClientConf
	Video zrpc.RpcClientConf
	File  zrpc.RpcClientConf
}
