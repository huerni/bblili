package logic

import (
	"context"

	"bblili/service/file/file"
	"bblili/service/file/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileMD5Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileMD5Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileMD5Logic {
	return &GetFileMD5Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileMD5Logic) GetFileMD5(in *file.GetFileMD5Request) (*file.GetFileMD5Response, error) {
	// todo: add your logic here and delete this line

	return &file.GetFileMD5Response{}, nil
}
