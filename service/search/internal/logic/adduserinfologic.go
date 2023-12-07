package logic

import (
	"context"
	"fmt"

	"bblili/service/search/internal/svc"
	"bblili/service/search/search"

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

func (l *AddUserInfoLogic) AddUserInfo(in *search.AddUserInfoRequest) (*search.AddUserInfoResponse, error) {
	// todo: add your logic here and delete this line
	indexName := "userinfos"
	exists, err := l.svcCtx.ElasClient.IndexExists(indexName).Do(l.ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		createIndex, err := l.svcCtx.ElasClient.CreateIndex(indexName).Do(l.ctx)
		if err != nil {
			return nil, err
		}
		if !createIndex.Acknowledged {
			return nil, fmt.Errorf("Create index not acknowledged")
		}
	}
	_, err = l.svcCtx.ElasClient.Index().
		Index(indexName).
		BodyJson(in.UserInfo).
		Do(l.ctx)
	if err != nil {
		return nil, err
	}
	return &search.AddUserInfoResponse{}, nil
}
