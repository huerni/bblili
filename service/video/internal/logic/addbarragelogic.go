package logic

import (
	"bblili/service/video/internal/db"
	"context"
	"gorm.io/gorm"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBarrageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBarrageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBarrageLogic {
	return &AddBarrageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加弹幕
func (l *AddBarrageLogic) AddBarrage(in *video.AddBarrageRequset) (*video.AddBarrageResponse, error) {
	if err := db.InsertDanMu(l.ctx, &db.Barrage{
		Model:       gorm.Model{},
		UserId:      in.BarrageInfo.UserId,
		VideoId:     in.BarrageInfo.VideoId,
		Content:     in.BarrageInfo.Content,
		BarrageTime: in.BarrageInfo.BarrageTime,
	}); err != nil {
		return nil, err
	}

	return &video.AddBarrageResponse{}, nil
}
