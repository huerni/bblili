package logic

import (
	"bblili/service/user/internal/common/constant"
	"bblili/service/user/internal/db"
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2/primitive"
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

	dbuserinfo := db.UserInfo{
		Model:  gorm.Model{},
		Nick:   in.UserInfo.Nick,
		Avatar: in.UserInfo.Avatar,
		Sign:   in.UserInfo.Sign,
		Gender: in.UserInfo.Gender,
		Birth:  in.UserInfo.Birth,
		UserId: in.UserInfo.UserId,
	}

	if err := db.InsertUserInfo(l.ctx, &dbuserinfo); err != nil {
		return nil, err
	}

	// 通过rocketmq通知到search服务
	str, err := json.Marshal(dbuserinfo)
	if err != nil {
		return nil, err
	}
	msg := &primitive.Message{
		Topic: constant.ROCKETMQ_ADDUSERINFO_TOPIC,
		Body:  str,
	}
	_, err = l.svcCtx.MQProducer.SendSync(l.ctx, msg)
	if err != nil {
		logx.Error("发送rocketmq消息失败")
		return nil, err
	}

	return &user.AddUserInfoResponse{}, nil
}
