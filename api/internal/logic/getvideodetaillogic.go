package logic

import (
	"bblili/service/user/userclient"
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

	userres, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &userclient.GetUserInfoRequest{UserId: res.Video.UserId})

	resp.UserInfo = &types.UserInfo{
		Id:     userres.UserInfo.Id,
		Nick:   userres.UserInfo.Nick,
		Avatar: userres.UserInfo.Avatar,
		Sign:   userres.UserInfo.Sign,
		Gender: userres.UserInfo.Gender,
		Birth:  userres.UserInfo.Birth,
	}

	return
}
