package myrocketmq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"log"
)

func RegisterRocketProducerMust(addrs []string, groupName string, retry int) rocketmq.Producer {
	res, err := rocketmq.NewProducer(
		producer.WithGroupName(groupName),
		producer.WithNsResolver(primitive.NewPassthroughResolver(addrs)),
		producer.WithRetry(retry),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = res.Start()
	if err != nil {
		log.Fatal(err)
	}

	return res
}
