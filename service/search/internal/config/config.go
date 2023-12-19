package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	ElasticUrl string         `json:",optional"`
	RocketMQ   RocketMQConfig `json:"rocketMQ"`
}

type RocketMQConfig struct {
	NameServer string `json:"nameServer"`
}
