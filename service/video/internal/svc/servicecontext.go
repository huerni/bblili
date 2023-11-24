package svc

import (
	"bblili/service/user/userclient"
	"bblili/service/video/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserClient: userclient.NewUser(zrpc.MustNewClient(c.User)),
	}
}
