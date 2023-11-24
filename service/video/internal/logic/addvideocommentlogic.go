package logic

import (
	"context"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoCommentLogic {
	return &AddVideoCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddVideoCommentLogic) AddVideoComment(in *video.AddVideoCommentRequest) (*video.AddVideoCommentResponse, error) {
	// todo: add your logic here and delete this line

	return &video.AddVideoCommentResponse{}, nil
}
