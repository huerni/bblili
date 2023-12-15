package logic

import (
	"bblili/service/video/videoclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperatorUserSanLianLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperatorUserSanLianLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperatorUserSanLianLogic {
	return &OperatorUserSanLianLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperatorUserSanLianLogic) OperatorUserSanLian(req *types.OperationSanLianRequest) (resp *types.OperationSanLianResponse, err error) {

	_, err = l.svcCtx.VideoClient.OperateVideoSanLian(l.ctx, &videoclient.OperateVideoSanLianRequest{
		VideoId:  req.VideoId,
		UserId:   req.UserId,
		GroupId:  req.GroupId,
		Operator: req.Operator,
	})

	return
}
