// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package videoclient

import (
	"context"

	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddBarrageRequset           = video.AddBarrageRequset
	AddBarrageResponse          = video.AddBarrageResponse
	AddVideoCommentRequest      = video.AddVideoCommentRequest
	AddVideoCommentResponse     = video.AddVideoCommentResponse
	AddVideoRequest             = video.AddVideoRequest
	AddVideoResponse            = video.AddVideoResponse
	BarrageInfo                 = video.BarrageInfo
	GetBarragesRequest          = video.GetBarragesRequest
	GetBarragesResponse         = video.GetBarragesResponse
	GetVideoCommentsRequest     = video.GetVideoCommentsRequest
	GetVideoCommentsResponse    = video.GetVideoCommentsResponse
	GetVideoDetailsRequest      = video.GetVideoDetailsRequest
	GetVideoDetailsResponse     = video.GetVideoDetailsResponse
	GetVideoSanLianRequest      = video.GetVideoSanLianRequest
	GetVideoSanLianResponse     = video.GetVideoSanLianResponse
	GetVideoUserSanLianRequest  = video.GetVideoUserSanLianRequest
	GetVideoUserSanLianResponse = video.GetVideoUserSanLianResponse
	GetVideosRequest            = video.GetVideosRequest
	GetVideosResponse           = video.GetVideosResponse
	OperateVideoSanLianRequest  = video.OperateVideoSanLianRequest
	OperateVideoSanLianResponse = video.OperateVideoSanLianResponse
	RecommendRequest            = video.RecommendRequest
	RecommendResponse           = video.RecommendResponse
	Request                     = video.Request
	Response                    = video.Response
	VideoComment                = video.VideoComment
	VideoInfo                   = video.VideoInfo
	ViewVideoRequest            = video.ViewVideoRequest
	ViewVideoResponse           = video.ViewVideoResponse

	Video interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		AddVideo(ctx context.Context, in *AddVideoRequest, opts ...grpc.CallOption) (*AddVideoResponse, error)
		GetVideos(ctx context.Context, in *GetVideosRequest, opts ...grpc.CallOption) (*GetVideosResponse, error)
		ViewVideo(ctx context.Context, in *ViewVideoRequest, opts ...grpc.CallOption) (*ViewVideoResponse, error)
		GetVideoDetails(ctx context.Context, in *GetVideoDetailsRequest, opts ...grpc.CallOption) (*GetVideoDetailsResponse, error)
		Recommend(ctx context.Context, in *RecommendRequest, opts ...grpc.CallOption) (*RecommendResponse, error)
		GetVideoSanLian(ctx context.Context, in *GetVideoSanLianRequest, opts ...grpc.CallOption) (*GetVideoSanLianResponse, error)
		GetVideoUserSanLian(ctx context.Context, in *GetVideoUserSanLianRequest, opts ...grpc.CallOption) (*GetVideoUserSanLianResponse, error)
		OperateVideoSanLian(ctx context.Context, in *OperateVideoSanLianRequest, opts ...grpc.CallOption) (*OperateVideoSanLianResponse, error)
		AddVideoComment(ctx context.Context, in *AddVideoCommentRequest, opts ...grpc.CallOption) (*AddVideoCommentResponse, error)
		GetVideoComments(ctx context.Context, in *GetVideoCommentsRequest, opts ...grpc.CallOption) (*GetVideoCommentsResponse, error)
		// 添加弹幕
		AddBarrage(ctx context.Context, in *AddBarrageRequset, opts ...grpc.CallOption) (*AddBarrageResponse, error)
		GetBarrages(ctx context.Context, in *GetBarragesRequest, opts ...grpc.CallOption) (*GetBarragesResponse, error)
	}

	defaultVideo struct {
		cli zrpc.Client
	}
)

func NewVideo(cli zrpc.Client) Video {
	return &defaultVideo{
		cli: cli,
	}
}

func (m *defaultVideo) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultVideo) AddVideo(ctx context.Context, in *AddVideoRequest, opts ...grpc.CallOption) (*AddVideoResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.AddVideo(ctx, in, opts...)
}

func (m *defaultVideo) GetVideos(ctx context.Context, in *GetVideosRequest, opts ...grpc.CallOption) (*GetVideosResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetVideos(ctx, in, opts...)
}

func (m *defaultVideo) ViewVideo(ctx context.Context, in *ViewVideoRequest, opts ...grpc.CallOption) (*ViewVideoResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.ViewVideo(ctx, in, opts...)
}

func (m *defaultVideo) GetVideoDetails(ctx context.Context, in *GetVideoDetailsRequest, opts ...grpc.CallOption) (*GetVideoDetailsResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetVideoDetails(ctx, in, opts...)
}

func (m *defaultVideo) Recommend(ctx context.Context, in *RecommendRequest, opts ...grpc.CallOption) (*RecommendResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.Recommend(ctx, in, opts...)
}

func (m *defaultVideo) GetVideoSanLian(ctx context.Context, in *GetVideoSanLianRequest, opts ...grpc.CallOption) (*GetVideoSanLianResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetVideoSanLian(ctx, in, opts...)
}

func (m *defaultVideo) GetVideoUserSanLian(ctx context.Context, in *GetVideoUserSanLianRequest, opts ...grpc.CallOption) (*GetVideoUserSanLianResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetVideoUserSanLian(ctx, in, opts...)
}

func (m *defaultVideo) OperateVideoSanLian(ctx context.Context, in *OperateVideoSanLianRequest, opts ...grpc.CallOption) (*OperateVideoSanLianResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.OperateVideoSanLian(ctx, in, opts...)
}

func (m *defaultVideo) AddVideoComment(ctx context.Context, in *AddVideoCommentRequest, opts ...grpc.CallOption) (*AddVideoCommentResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.AddVideoComment(ctx, in, opts...)
}

func (m *defaultVideo) GetVideoComments(ctx context.Context, in *GetVideoCommentsRequest, opts ...grpc.CallOption) (*GetVideoCommentsResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetVideoComments(ctx, in, opts...)
}

// 添加弹幕
func (m *defaultVideo) AddBarrage(ctx context.Context, in *AddBarrageRequset, opts ...grpc.CallOption) (*AddBarrageResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.AddBarrage(ctx, in, opts...)
}

func (m *defaultVideo) GetBarrages(ctx context.Context, in *GetBarragesRequest, opts ...grpc.CallOption) (*GetBarragesResponse, error) {
	client := video.NewVideoClient(m.cli.Conn())
	return client.GetBarrages(ctx, in, opts...)
}
