// Code generated by goctl. DO NOT EDIT.
// Source: file.proto

package fileclient

import (
	"context"

	"bblili/service/file/file"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeleteFileRequest         = file.DeleteFileRequest
	DeleteFileResponse        = file.DeleteFileResponse
	GetFileMD5Request         = file.GetFileMD5Request
	GetFileMD5Response        = file.GetFileMD5Response
	GetFileUrlByMD5Request    = file.GetFileUrlByMD5Request
	GetFileUrlByMD5Response   = file.GetFileUrlByMD5Response
	GetSuccessChunkRequest    = file.GetSuccessChunkRequest
	GetSuccessChunkResponse   = file.GetSuccessChunkResponse
	Request                   = file.Request
	Response                  = file.Response
	UpLoadFileBySlicesRequest = file.UpLoadFileBySlicesRequest
	UploadFileBySliceResponse = file.UploadFileBySliceResponse

	File interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		UploadFileSlices(ctx context.Context, in *UpLoadFileBySlicesRequest, opts ...grpc.CallOption) (*UploadFileBySliceResponse, error)
		DeleteFileMD5(ctx context.Context, in *GetFileMD5Request, opts ...grpc.CallOption) (*GetFileMD5Response, error)
		GetFileMD5(ctx context.Context, in *GetFileMD5Request, opts ...grpc.CallOption) (*GetFileMD5Response, error)
		GetSuccessChunk(ctx context.Context, in *GetSuccessChunkRequest, opts ...grpc.CallOption) (*GetSuccessChunkResponse, error)
		GetFileUrlByMD5(ctx context.Context, in *GetFileUrlByMD5Request, opts ...grpc.CallOption) (*GetFileUrlByMD5Response, error)
	}

	defaultFile struct {
		cli zrpc.Client
	}
)

func NewFile(cli zrpc.Client) File {
	return &defaultFile{
		cli: cli,
	}
}

func (m *defaultFile) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := file.NewFileClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultFile) UploadFileSlices(ctx context.Context, in *UpLoadFileBySlicesRequest, opts ...grpc.CallOption) (*UploadFileBySliceResponse, error) {
	client := file.NewFileClient(m.cli.Conn())
	return client.UploadFileSlices(ctx, in, opts...)
}

func (m *defaultFile) DeleteFileMD5(ctx context.Context, in *GetFileMD5Request, opts ...grpc.CallOption) (*GetFileMD5Response, error) {
	client := file.NewFileClient(m.cli.Conn())
	return client.DeleteFileMD5(ctx, in, opts...)
}

func (m *defaultFile) GetFileMD5(ctx context.Context, in *GetFileMD5Request, opts ...grpc.CallOption) (*GetFileMD5Response, error) {
	client := file.NewFileClient(m.cli.Conn())
	return client.GetFileMD5(ctx, in, opts...)
}

func (m *defaultFile) GetSuccessChunk(ctx context.Context, in *GetSuccessChunkRequest, opts ...grpc.CallOption) (*GetSuccessChunkResponse, error) {
	client := file.NewFileClient(m.cli.Conn())
	return client.GetSuccessChunk(ctx, in, opts...)
}

func (m *defaultFile) GetFileUrlByMD5(ctx context.Context, in *GetFileUrlByMD5Request, opts ...grpc.CallOption) (*GetFileUrlByMD5Response, error) {
	client := file.NewFileClient(m.cli.Conn())
	return client.GetFileUrlByMD5(ctx, in, opts...)
}
