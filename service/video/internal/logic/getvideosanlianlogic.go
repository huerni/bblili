package logic

import (
	"bblili/service/video/internal/db"
	"context"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoSanLianLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoSanLianLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoSanLianLogic {
	return &GetVideoSanLianLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoSanLianLogic) GetVideoSanLian(in *video.GetVideoSanLianRequest) (*video.GetVideoSanLianResponse, error) {
	coinCount, err := db.CountVideoCoinByVideoId(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}

	collectionCount, err := db.CountVideoCollectionByVideoId(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}

	likeCount, err := db.CountVideoLikeByVideoId(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}

	return &video.GetVideoSanLianResponse{CoinCount: int32(*coinCount), CollectionCount: int32(*collectionCount), LikeCount: int32(*likeCount)}, nil
}
