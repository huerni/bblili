package logic

import (
	"bblili/service/user/userclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {

	_, err = l.svcCtx.UserClient.Login(l.ctx, &userclient.LoginRequest{
		UsernameOrPhoneOrEmail: req.UsernameOrPhoneOrEmail,
		Password:               req.Password,
	})
	if err != nil {
		return
	}
	resp.AuthToken = "1"
	return
}
