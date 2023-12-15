package svc

import (
	"bblili/api/internal/config"
	"bblili/service/file/fileclient"
	"bblili/service/user/userclient"
	"bblili/service/video/videoclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserClient  userclient.User
	VideoClient videoclient.Video
	FileClient  fileclient.File
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserClient:  userclient.NewUser(zrpc.MustNewClient(c.User)),
		VideoClient: videoclient.NewVideo(zrpc.MustNewClient(c.Video)),
		FileClient:  fileclient.NewFile(zrpc.MustNewClient(c.File)),
	}
}
