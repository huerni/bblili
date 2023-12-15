package logic

import (
	"bblili/service/video/videoclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoCommentLogic {
	return &AddVideoCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddVideoCommentLogic) AddVideoComment(req *types.AddVideoCommentRequest) (resp *types.AddVideoCommentResponse, err error) {

	_, err = l.svcCtx.VideoClient.AddVideoComment(l.ctx, &videoclient.AddVideoCommentRequest{
		VideoId:     req.VideoId,
		UserId:      req.UserId,
		Comment:     req.Comment,
		ReplyUserId: req.ReplyUserId,
		RootId:      req.RootId,
	})

	return
}
