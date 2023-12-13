package logic

import (
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
	// todo: add your logic here and delete this line

	return
}
