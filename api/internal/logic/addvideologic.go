package logic

import (
	"bblili/service/video/videoclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoLogic {
	return &AddVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddVideoLogic) AddVideo(req *types.AddVideoRequest) (resp *types.AddVideoResponse, err error) {

	_, err = l.svcCtx.VideoClient.AddVideo(l.ctx, &videoclient.AddVideoRequest{
		UserId:      req.UserId,
		Thumbnail:   req.Thumbnail,
		Title:       req.Title,
		Types:       req.Types,
		Duration:    req.Duration,
		Area:        req.Area,
		Description: req.Description,
		TagList:     req.VideoTagList,
		FileMd5:     req.FileMd5,
	})

	return
}
