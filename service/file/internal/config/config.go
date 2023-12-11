package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DSN   string    `json:"dsn"`
	MinIO MinIOConf `json:"minIO"`
}

type MinIOConf struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
	UseSSL          bool   `json:"useSSL,default=false"`
}

var MINIO_BUCKET = "bblili"
var MINIO_BASE_PATH = "breakpoint"
var MINIO_LOCATION = "cn-north-1"
