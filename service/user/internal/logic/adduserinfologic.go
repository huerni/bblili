package logic

import (
	"bblili/service/user/internal/db"
	"context"
	"fmt"
	"gorm.io/gorm"

	"bblili/service/user/internal/svc"
	"bblili/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserInfoLogic {
	return &AddUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserInfoLogic) AddUserInfo(in *user.AddUserInfoRequest) (*user.AddUserInfoResponse, error) {
	queryUser, err := db.QueryUserById(l.ctx, in.GetUserInfo().UserId)
	if err != nil {
		return nil, err
	}
	if queryUser == nil {
		return nil, fmt.Errorf("该用户不存在!")
	}
	if err := db.InsertUserInfo(l.ctx, &db.UserInfo{
		Model:  gorm.Model{},
		Nick:   in.UserInfo.Nick,
		Avatar: in.UserInfo.Avatar,
		Sign:   in.UserInfo.Sign,
		Gender: in.UserInfo.Gender,
		Birth:  in.UserInfo.Birth,
		UserId: in.UserInfo.UserId,
	}); err != nil {
		return nil, err
	}
	return &user.AddUserInfoResponse{}, nil
}
