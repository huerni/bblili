package logic

import (
	"bblili/service/user/internal/db"
	"context"

	"bblili/service/user/internal/svc"
	"bblili/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	queryUser, err := db.QueryUserById(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return &user.GetUserResponse{Username: queryUser.Username, Password: queryUser.Password, Phone: queryUser.Phone, Email: queryUser.Email}, nil
}
