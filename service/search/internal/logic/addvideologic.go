package logic

import (
	"bblili/service/search/internal/svc"
	"bblili/service/search/search"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoLogic {
	return &AddVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddVideoLogic) AddVideo(in *search.AddVideoRequest) (*search.AddVideoResponse, error) {
	indexName := "videos"
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
	_, err = l.svcCtx.ElasClient.Index().Index(indexName).BodyJson(in.Video).Do(l.ctx)
	if err != nil {
		return nil, err
	}
	return &search.AddVideoResponse{}, nil
}
