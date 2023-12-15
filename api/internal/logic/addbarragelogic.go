package logic

import (
	"bblili/service/video/videoclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBarrageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBarrageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBarrageLogic {
	return &AddBarrageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBarrageLogic) AddBarrage(req *types.AddBarrageRequest) (resp *types.AddBarrageResponse, err error) {

	_, err = l.svcCtx.VideoClient.AddBarrage(l.ctx, &videoclient.AddBarrageRequset{BarrageInfo: &videoclient.BarrageInfo{UserId: req.BarrageInfo.UserId, VideoId: req.BarrageInfo.UserId, Content: req.BarrageInfo.Content, BarrageTime: req.BarrageInfo.BarrageTime}})

	return
}
