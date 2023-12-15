package logic

import (
	"bblili/service/video/internal/db"
	"context"
	"gorm.io/gorm"

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

	if err := db.InsertVideoComment(l.ctx, &db.VideoComment{
		Model:       gorm.Model{},
		VideoId:     in.VideoId,
		UserId:      in.UserId,
		Comment:     in.Comment,
		ReplyUserId: in.ReplyUserId,
		RootId:      in.RootId,
	}); err != nil {
		return nil, err
	}

	return &video.AddVideoCommentResponse{}, nil
}
