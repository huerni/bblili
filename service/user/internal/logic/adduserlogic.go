package logic

import (
	"bblili/service/user/internal/db"
	"context"
	"fmt"

	"bblili/service/user/internal/svc"
	"bblili/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.AddUserRequest) (*user.AddUserResponse, error) {
	// todo: add your logic here and delete this line

	users, err := db.QueryUser(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if users != nil {
		return nil, fmt.Errorf("用户名已存在")
	}
	users, err = db.QueryUserByPhone(l.ctx, in.Phone)
	if err != nil {
		return nil, err
	}
	if users != nil {
		return nil, fmt.Errorf("电话号码已存在")
	}
	users, err = db.QueryUserByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if users != nil {
		return nil, fmt.Errorf("邮箱已存在")
	}
	var insertUser = db.User{
		Username: in.Username,
		Phone:    in.Phone,
		Email:    in.Email,
		Password: in.Password,
	}
	if err = db.InsertUser(l.ctx, &insertUser); err != nil {
		return nil, err
	}

	return &user.AddUserResponse{}, nil
}
