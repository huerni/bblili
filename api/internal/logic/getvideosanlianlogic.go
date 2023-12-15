package logic

import (
	"bblili/service/video/videoclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoSanLianLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoSanLianLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoSanLianLogic {
	return &GetVideoSanLianLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoSanLianLogic) GetVideoSanLian(req *types.GetVideoSanLianRequest) (resp *types.GetVideoSanLianResponse, err error) {

	res, err := l.svcCtx.VideoClient.GetVideoSanLian(l.ctx, &videoclient.GetVideoSanLianRequest{VideoId: req.VideoId})

	if err != nil {
		return
	}

	resp.CoinCount = res.CoinCount
	resp.LikeCount = res.LikeCount
	resp.CollectionCount = res.CollectionCount

	return
}
