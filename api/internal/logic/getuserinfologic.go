package logic

import (
	"bblili/api/internal/svc"
	"bblili/api/internal/types"
	"bblili/service/user/userclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {

	res, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &userclient.GetUserInfoRequest{UserId: req.UserId})

	if err != nil {
		return
	}

	resp.UserInfo = types.UserInfo{
		Id:     res.UserInfo.Id,
		Nick:   res.UserInfo.Nick,
		Avatar: res.UserInfo.Avatar,
		Sign:   res.UserInfo.Sign,
		Gender: res.UserInfo.Gender,
		Birth:  res.UserInfo.Birth,
	}

	return
}
