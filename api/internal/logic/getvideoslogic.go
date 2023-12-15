package logic

import (
	"bblili/service/video/videoclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideosLogic {
	return &GetVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideosLogic) GetVideos(req *types.GetVideosRequest) (resp *types.GetVideosResponse, err error) {

	res, err := l.svcCtx.VideoClient.GetVideos(l.ctx, &videoclient.GetVideosRequest{
		Size: req.Size,
		Page: req.Page,
		Area: req.Area,
	})

	resp.VideoInfos = make([]*types.VideoInfo, 0)

	for _, videoinfo := range res.Videos {
		resp.VideoInfos = append(resp.VideoInfos, &types.VideoInfo{
			UserId:       videoinfo.UserId,
			Url:          videoinfo.Url,
			Thumbnail:    videoinfo.Thumbnail,
			Title:        videoinfo.Title,
			Types:        videoinfo.Types,
			Duration:     videoinfo.Duration,
			Area:         videoinfo.Area,
			Description:  videoinfo.Description,
			VideoTagList: videoinfo.TagList,
		})
	}

	return
}
