package logic

import (
	"bblili/service/video/videoclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoDetailLogic {
	return &GetVideoDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoDetailLogic) GetVideoDetail(req *types.GetVideoDetailRequest) (resp *types.GetVideoDetailResponse, err error) {
	res, err := l.svcCtx.VideoClient.GetVideoDetails(l.ctx, &videoclient.GetVideoDetailsRequest{VideoId: req.VideoId})

	if err != nil {
		return
	}

	resp.VideoInfo = &types.VideoInfo{
		UserId:       res.Video.UserId,
		Url:          res.Video.Url,
		Thumbnail:    res.Video.Thumbnail,
		Title:        res.Video.Title,
		Types:        res.Video.Types,
		Duration:     res.Video.Duration,
		Area:         res.Video.Area,
		Description:  res.Video.Description,
		VideoTagList: res.Video.TagList,
	}

	resp.UserInfo = &types.UserInfo{
		Id:     res.UserInfo.Id,
		Nick:   res.UserInfo.Nick,
		Avatar: res.UserInfo.Avatar,
		Sign:   res.UserInfo.Sign,
		Gender: res.UserInfo.Gender,
		Birth:  res.UserInfo.Birth,
	}

	return
}
