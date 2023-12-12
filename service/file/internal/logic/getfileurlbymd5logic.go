package logic

import (
	"bblili/service/file/internal/db"
	"context"

	"bblili/service/file/file"
	"bblili/service/file/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileUrlByMD5Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileUrlByMD5Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileUrlByMD5Logic {
	return &GetFileUrlByMD5Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileUrlByMD5Logic) GetFileUrlByMD5(in *file.GetFileUrlByMD5Request) (*file.GetFileUrlByMD5Response, error) {
	qfile, err := db.QueryFileByMD5(l.ctx, in.Md5)
	if err != nil {
		return nil, err
	}

	return &file.GetFileUrlByMD5Response{Url: qfile.Url}, nil
}
