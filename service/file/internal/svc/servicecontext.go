package svc

import (
	"bblili/service/file/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type ServiceContext struct {
	Config      config.Config
	MinIOClient *minio.Client
	MinIOCore   *minio.Core
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := minio.New(c.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.MinIO.AccessKeyID, c.MinIO.SecretAccessKey, ""),
		Secure: c.MinIO.UseSSL,
	})
	if err != nil {
		log.Println("minio连接错误", err)
	}
	core, err := minio.NewCore(c.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.MinIO.AccessKeyID, c.MinIO.SecretAccessKey, ""),
		Secure: c.MinIO.UseSSL,
	})
	if err != nil {
		log.Println("minio连接错误", err)
	}
	return &ServiceContext{
		Config:      c,
		MinIOClient: client,
		MinIOCore:   core,
	}
}
