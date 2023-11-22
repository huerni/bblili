package logic

import (
	"bblili/service/user/internal/db"
	"context"
	"fmt"

	"bblili/service/user/internal/svc"
	"bblili/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	var users []*db.User
	var err error

	if users, err = db.QueryUser(l.ctx, in.UsernameOrPhoneOrEmail); err != nil {
		return nil, err
	}
	if users != nil && users[0].Password != in.Password {
		return nil, fmt.Errorf("账号密码不正确")
	} else if users == nil {
		if users, err = db.QueryUserByPhone(l.ctx, in.UsernameOrPhoneOrEmail); err != nil {
			return nil, err
		}
		if users != nil && users[0].Password != in.Password {
			return nil, fmt.Errorf("账号密码不正确")
		} else if users == nil {
			if users, err = db.QueryUserByEmail(l.ctx, in.UsernameOrPhoneOrEmail); err != nil {
				return nil, err
			}
			if users != nil && users[0].Password != in.Password {
				return nil, fmt.Errorf("账号密码不正确")
			} else if users == nil {
				return nil, fmt.Errorf("账号不存在")
			}
		}
	}
	return &user.LoginResponse{}, nil
}
