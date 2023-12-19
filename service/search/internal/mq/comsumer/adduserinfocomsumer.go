package comsumer

import (
	myrocketmq "bblili/pkg/rocketmq"
	"bblili/service/search/internal/config"
	"bblili/service/search/internal/constant"
	"bblili/service/search/internal/svc"
	"bblili/service/search/search"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

func NewAddUserInfoComsumer(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) {
	client := myrocketmq.RegisterComsumer([]string{c.RocketMQ.NameServer}, constant.BBLILI_ROCKETMQ_GROUP)
	err := client.Subscribe(constant.ROCKETMQ_ADDUSERINFO_TOPIC, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		indexName := "userinfos"
		exists, err := svcCtx.ElasClient.IndexExists(indexName).Do(ctx)
		if err != nil {
			return consumer.ConsumeRetryLater, err
		}
		if !exists {
			createIndex, err := svcCtx.ElasClient.CreateIndex(indexName).Do(ctx)
			if err != nil {
				return consumer.ConsumeRetryLater, err
			}
			if !createIndex.Acknowledged {
				return consumer.ConsumeRetryLater, fmt.Errorf("Create index not acknowledged")
			}
		}
		var userinfo search.UserInfo
		err = json.Unmarshal(msgs[0].Body, &userinfo)
		if err != nil {
			log.Fatal(err)
		}
		_, err = svcCtx.ElasClient.Index().
			Index(indexName).
			BodyJson(&userinfo).
			Do(ctx)
		if err != nil {
			return consumer.ConsumeRetryLater, err
		}
		logx.Info("消费数据成功！！")
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		logx.Error(err)
	}

	// 开启消费
	err = client.Start()
	if err != nil {
		logx.Error(err)
	}
}
