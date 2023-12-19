package logic

import (
	"bblili/service/video/internal/db"
	"context"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideosLogic {
	return &GetVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideosLogic) GetVideos(in *video.GetVideosRequest) (*video.GetVideosResponse, error) {
	var qvideo []*db.Video
	var err error

	if in.Area != 0 {
		qvideo, err = db.QueryVideoByArea(l.ctx, in.Page, in.Size, in.Area)
		if err != nil {
			return nil, err
		}
	}

	qvideo, err = db.QueryVideo(l.ctx, in.Page, in.Size)
	if err != nil {
		return nil, err
	}

	videoInfos := make([]*video.VideoInfo, 0)
	for _, videoinfo := range qvideo {
		videoInfos = append(videoInfos, &video.VideoInfo{
			Id:          uint64(videoinfo.ID),
			UserId:      videoinfo.UserId,
			Url:         videoinfo.Url,
			Thumbnail:   videoinfo.Thumbnail,
			Title:       videoinfo.Title,
			Types:       videoinfo.Types,
			Duration:    videoinfo.Duration,
			Area:        videoinfo.Area,
			Description: videoinfo.Description,
			TagList:     nil,
		})
	}

	return &video.GetVideosResponse{Videos: videoInfos}, nil
}
