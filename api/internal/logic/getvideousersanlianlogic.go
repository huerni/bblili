package logic

import (
	"bblili/service/video/videoclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoUserSanLianLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoUserSanLianLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoUserSanLianLogic {
	return &GetVideoUserSanLianLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoUserSanLianLogic) GetVideoUserSanLian(req *types.GetVideoUserSanLianRequest) (resp *types.GetVideoUserSanLianResponse, err error) {

	res, err := l.svcCtx.VideoClient.GetVideoUserSanLian(l.ctx, &videoclient.GetVideoUserSanLianRequest{
		VideoId: req.VideoId,
		UserId:  req.UserId,
	})

	if err != nil {
		return
	}

	resp.CoinCount = res.CoinCount
	resp.IsLike = res.IsLike
	resp.IsCollection = res.IsCollection

	return
}
