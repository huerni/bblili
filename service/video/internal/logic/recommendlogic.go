package logic

import (
	"context"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecommendLogic) Recommend(in *video.RecommendRequest) (*video.RecommendResponse, error) {
	// todo: add your logic here and delete this line

	return &video.RecommendResponse{}, nil
}
