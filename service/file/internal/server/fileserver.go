// Code generated by goctl. DO NOT EDIT.
// Source: file.proto

package server

import (
	"context"

	"bblili/service/file/file"
	"bblili/service/file/internal/logic"
	"bblili/service/file/internal/svc"
)

type FileServer struct {
	svcCtx *svc.ServiceContext
	file.UnimplementedFileServer
}

func NewFileServer(svcCtx *svc.ServiceContext) *FileServer {
	return &FileServer{
		svcCtx: svcCtx,
	}
}

func (s *FileServer) Ping(ctx context.Context, in *file.Request) (*file.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *FileServer) UploadFileSlices(ctx context.Context, in *file.UpLoadFileBySlicesRequest) (*file.UploadFileBySliceResponse, error) {
	l := logic.NewUploadFileSlicesLogic(ctx, s.svcCtx)
	return l.UploadFileSlices(in)
}

func (s *FileServer) DeleteFileMD5(ctx context.Context, in *file.GetFileMD5Request) (*file.GetFileMD5Response, error) {
	l := logic.NewDeleteFileMD5Logic(ctx, s.svcCtx)
	return l.DeleteFileMD5(in)
}

func (s *FileServer) GetFileMD5(ctx context.Context, in *file.GetFileMD5Request) (*file.GetFileMD5Response, error) {
	l := logic.NewGetFileMD5Logic(ctx, s.svcCtx)
	return l.GetFileMD5(in)
}

func (s *FileServer) GetSuccessChunk(ctx context.Context, in *file.GetSuccessChunkRequest) (*file.GetSuccessChunkResponse, error) {
	l := logic.NewGetSuccessChunkLogic(ctx, s.svcCtx)
	return l.GetSuccessChunk(in)
}

func (s *FileServer) GetFileUrlByMD5(ctx context.Context, in *file.GetFileUrlByMD5Request) (*file.GetFileUrlByMD5Response, error) {
	l := logic.NewGetFileUrlByMD5Logic(ctx, s.svcCtx)
	return l.GetFileUrlByMD5(in)
}
