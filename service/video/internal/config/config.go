package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DSN      string `json:"dsn"`
	File     zrpc.RpcClientConf
	RocketMQ RocketMQConfig `json:"rocketMQ"`
}

type RocketMQConfig struct {
	NameServer string `json:"nameServer"`
}
