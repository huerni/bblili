package logic

import (
	"bblili/service/file/fileclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSuccessChunkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSuccessChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSuccessChunkLogic {
	return &GetSuccessChunkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSuccessChunkLogic) GetSuccessChunk(req *types.GetSuccessChunkRequest) (resp *types.GetSuccessChunkResponse, err error) {

	res, err := l.svcCtx.FileClient.GetSuccessChunk(l.ctx, &fileclient.GetSuccessChunkRequest{FileMD5: req.FileMd5, Filetype: req.Filetype, TotalChunkCount: req.TotalChunkCount})
	if err != nil {
		return
	}

	resp.UploadID = res.UploadID
	resp.UuId = res.UuId
	resp.Chunks = res.Chunks

	return
}
