package svc

import (
	myrocketmq "bblili/pkg/rocketmq"
	"bblili/service/file/fileclient"
	"bblili/service/video/internal/common/constant"
	"bblili/service/video/internal/config"
	"github.com/apache/rocketmq-client-go/v2"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	FileClient  fileclient.File
	RedisClient redis.Redis
	RedisSync   redsync.Redsync
	MQProducer  rocketmq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		FileClient:  fileclient.NewFile(zrpc.MustNewClient(c.File)),
		RedisClient: *redis.MustNewRedis(c.Redis.RedisConf),
		RedisSync: *redsync.New(goredis.NewPool(goredislib.NewClient(&goredislib.Options{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}))),
		MQProducer: myrocketmq.RegisterRocketProducerMust([]string{c.RocketMQ.NameServer}, constant.BBLILI_ROCKETMQ_GROUP, 0),
	}
}
