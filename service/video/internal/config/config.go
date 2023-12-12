package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DSN  string `json:"dsn"`
	User zrpc.RpcClientConf
	File zrpc.RpcClientConf
}
