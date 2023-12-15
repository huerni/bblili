package logic

import (
	"bblili/service/file/fileclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileChunksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileChunksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileChunksLogic {
	return &UploadFileChunksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileChunksLogic) UploadFileChunks(req *types.UploadFileChunksRequest) (resp *types.UploadFileChunksResponse, err error) {
	res, err := l.svcCtx.FileClient.UploadFileSlices(l.ctx, &fileclient.UpLoadFileBySlicesRequest{FileMD5: req.FileMd5, FileType: req.Filetype, SliceNumber: req.ChunkNumber, SliceSize: req.ChunkSize})
	if err != nil {
		return
	}

	resp.Url = res.Url

	return
}
