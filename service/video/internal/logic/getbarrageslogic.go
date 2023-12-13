package logic

import (
	"bblili/service/video/internal/db"
	"context"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBarragesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBarragesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBarragesLogic {
	return &GetBarragesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBarragesLogic) GetBarrages(in *video.GetBarragesRequest) (*video.GetBarragesResponse, error) {
	barrages, err := db.QueryBarrageByVideoIdAndCurrentTime(l.ctx, in.VideoId, in.CurrentTime)
	if err != nil {
		return nil, err
	}

	barrageInfos := make([]*video.BarrageInfo, 0)
	for _, barrage := range barrages {
		barrageInfos = append(barrageInfos, &video.BarrageInfo{
			UserId:      barrage.UserId,
			VideoId:     barrage.VideoId,
			Content:     barrage.Content,
			BarrageTime: barrage.BarrageTime,
		})
	}

	return &video.GetBarragesResponse{BarrageInfos: barrageInfos}, nil
}
