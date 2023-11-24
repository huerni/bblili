package logic

import (
	"bblili/service/user/internal/db"
	"context"

	"bblili/service/user/internal/svc"
	"bblili/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	userInfo, err := db.QueryUserInfoByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return &user.GetUserInfoResponse{UserInfo: &user.UserInfo{
		Id:     uint64(userInfo.ID),
		Nick:   userInfo.Nick,
		Avatar: userInfo.Avatar,
		Sign:   userInfo.Sign,
		Gender: userInfo.Gender,
		Birth:  userInfo.Birth,
		UserId: userInfo.UserId,
	}}, nil
}
