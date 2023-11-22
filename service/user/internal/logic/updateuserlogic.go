package logic

import (
	"bblili/service/user/internal/db"
	"context"

	"bblili/service/user/internal/svc"
	"bblili/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	// todo: add your logic here and delete this line
	var err error
	userInfo := db.UserInfo{
		Nick:   in.GetUserInfo().Nick,
		Avatar: in.GetUserInfo().Avatar,
		Sign:   in.GetUserInfo().Sign,
		Gender: in.GetUserInfo().Gender,
		Birth:  in.GetUserInfo().Birth,
	}
	if err = db.UpdateUserInfo(l.ctx, &userInfo); err != nil {
		return nil, err
	}
	return &user.UpdateUserResponse{}, nil
}
