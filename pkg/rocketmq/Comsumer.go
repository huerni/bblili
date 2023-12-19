package myrocketmq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
)

func RegisterComsumer(addrs []string, groupName string) rocketmq.PushConsumer {
	res, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(groupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(addrs)),
	)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
