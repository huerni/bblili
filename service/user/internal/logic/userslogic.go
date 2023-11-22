package logic

import (
	"context"

	"bblili/service/user/internal/svc"
	"bblili/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersLogic {
	return &UsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UsersLogic) Users(in *user.UsersRequest) (*user.UsersResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UsersResponse{}, nil
}
