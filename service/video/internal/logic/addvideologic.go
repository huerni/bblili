package logic

import (
	"bblili/service/file/file"
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
	// 访问file服务获得url
	response, err := l.svcCtx.FileClient.GetFileUrlByMD5(l.ctx, &file.GetFileUrlByMD5Request{Md5: in.FileMd5})
	if err != nil {
		return nil, err
	}
	if err := db.InsertVideo(l.ctx, &db.Video{
		Model:     gorm.Model{},
		UserId:    in.UserId,
		Url:       response.Url,
		Thumbnail: in.Thumbnail,
		Title:     in.Title,
		Types:     in.Types,
		Duration:  in.Duration,
		Area:      in.Area,
	}); err != nil {
		return nil, err
	}
	return &video.AddVideoResponse{}, nil
}
