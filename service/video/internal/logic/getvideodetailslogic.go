package logic

import (
	"bblili/service/video/internal/db"
	"context"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoDetailsLogic {
	return &GetVideoDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoDetailsLogic) GetVideoDetails(in *video.GetVideoDetailsRequest) (*video.GetVideoDetailsResponse, error) {
	queryVideo, err := db.QueryVideoByVideoId(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}

	return &video.GetVideoDetailsResponse{Video: &video.VideoInfo{
		Id:          uint64(queryVideo.ID),
		UserId:      queryVideo.UserId,
		Url:         queryVideo.Url,
		Thumbnail:   queryVideo.Thumbnail,
		Title:       queryVideo.Title,
		Types:       queryVideo.Types,
		Duration:    queryVideo.Duration,
		Area:        queryVideo.Area,
		Description: queryVideo.Description,
		TagList:     nil,
	}}, nil
}
