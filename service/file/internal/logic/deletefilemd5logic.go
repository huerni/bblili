package logic

import (
	"context"

	"bblili/service/file/file"
	"bblili/service/file/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFileMD5Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFileMD5Logic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFileMD5Logic {
	return &DeleteFileMD5Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFileMD5Logic) DeleteFileMD5(in *file.GetFileMD5Request) (*file.GetFileMD5Response, error) {
	// todo: add your logic here and delete this line

	return &file.GetFileMD5Response{}, nil
}
