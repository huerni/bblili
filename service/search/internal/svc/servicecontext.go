package svc

import (
	"bblili/service/search/internal/config"
	"log"
)
import "github.com/olivere/elastic/v7"

type ServiceContext struct {
	Config     config.Config
	ElasClient *elastic.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	elasticClient, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(c.ElasticUrl))
	if err != nil || elasticClient == nil {
		log.Println("连接elastic失败", err)
	}

	return &ServiceContext{
		Config:     c,
		ElasClient: elasticClient,
	}
}
