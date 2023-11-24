package logic

import (
	"bblili/service/video/internal/db"
	"bblili/service/video/internal/svc"
	"bblili/service/video/video"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoUserSanLianLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoUserSanLianLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoUserSanLianLogic {
	return &GetVideoUserSanLianLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoUserSanLianLogic) GetVideoUserSanLian(in *video.GetVideoUserSanLianRequest) (*video.GetVideoUserSanLianResponse, error) {
	coins, err := db.QueryVideoCoinByVideoIdAndUserId(l.ctx, in.VideoId, in.UserId)
	if err != nil {
		return nil, err
	}
	coinCount := 0
	if coins != nil && len(coins) > 0 {
		coinCount = coins[0].Amount
	}
	likes, err := db.QueryVideoLikeByUserIdAndVideoId(l.ctx, in.UserId, in.VideoId)
	if err != nil {
		return nil, err
	}
	isLike := false
	if likes != nil && len(likes) > 0 {
		isLike = true
	}

	isCollection := false
	collections, err := db.QueryVideoCollectionByVideoIdAdnUserId(l.ctx, in.VideoId, in.UserId)
	if err != nil {
		return nil, err
	}
	if collections != nil && len(collections) > 0 {
		isCollection = true
	}

	return &video.GetVideoUserSanLianResponse{CoinCount: int32(coinCount), IsLike: isLike, IsCollection: isCollection}, nil
}
