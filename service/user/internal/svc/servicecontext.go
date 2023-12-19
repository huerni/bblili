package svc

import (
	myrocketmq "bblili/pkg/rocketmq"
	"bblili/service/user/internal/common/constant"
	"bblili/service/user/internal/config"
	"github.com/apache/rocketmq-client-go/v2"
)

type ServiceContext struct {
	Config     config.Config
	MQProducer rocketmq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MQProducer: myrocketmq.RegisterRocketProducerMust([]string{c.RocketMQ.NameServer}, constant.BBLILI_ROCKETMQ_GROUP, 0),
	}
}
