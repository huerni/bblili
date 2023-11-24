package logic

import (
	"bblili/service/video/internal/db"
	"context"
	"gorm.io/gorm"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

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

func (l *AddVideoLogic) AddVideo(in *video.AddVideoRequest) (*video.AddVideoResponse, error) {
	if err := db.InsertVideo(l.ctx, &db.Video{
		Model:     gorm.Model{},
		UserId:    in.GetVideo().UserId,
		Url:       in.GetVideo().Url,
		Thumbnail: in.GetVideo().Thumbnail,
		Title:     in.GetVideo().Title,
		Types:     in.GetVideo().Types,
		Duration:  in.GetVideo().Duration,
		Area:      in.GetVideo().Area,
	}); err != nil {
		return nil, err
	}
	return &video.AddVideoResponse{}, nil
}
