package logic

import (
	"bblili/service/video/internal/db"
	"bblili/service/video/internal/svc"
	"bblili/service/video/video"
	"context"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDanMuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDanMuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDanMuLogic {
	return &AddDanMuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 弹幕
func (l *AddDanMuLogic) AddDanMu(in *video.AddDanMuRequset) (*video.AddDanMuResponse, error) {
	if err := db.InsertDanMu(l.ctx, &db.DanMu{
		Model:     gorm.Model{},
		UserId:    in.DanmuInfo.UserId,
		VideoId:   in.DanmuInfo.VideoId,
		Content:   in.DanmuInfo.Content,
		DanmuTime: in.DanmuInfo.DanmuTime,
	}); err != nil {
		return nil, err
	}

	return &video.AddDanMuResponse{}, nil
}
