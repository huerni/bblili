package logic

import (
	"context"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewVideoLogic {
	return &ViewVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ViewVideoLogic) ViewVideo(in *video.ViewVideoRequest) (*video.ViewVideoResponse, error) {

	return &video.ViewVideoResponse{}, nil
}
