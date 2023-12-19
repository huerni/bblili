package logic

import (
	"bblili/service/file/file"
	"bblili/service/video/internal/common/constant"
	"bblili/service/video/internal/db"
	"bblili/service/video/internal/svc"
	"bblili/service/video/video"
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type AddVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoLogic {
	return &AddVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddVideoLogic) AddVideo(in *video.AddVideoRequest) (*video.AddVideoResponse, error) {
	// 访问file服务获得url
	response, err := l.svcCtx.FileClient.GetFileUrlByMD5(l.ctx, &file.GetFileUrlByMD5Request{Md5: in.FileMd5})
	if err != nil {
		logx.Error("远程调用file服务失败")
		return nil, err
	}

	dbvideo := db.Video{
		Model:     gorm.Model{},
		UserId:    in.UserId,
		Url:       response.Url,
		Thumbnail: in.Thumbnail,
		Title:     in.Title,
		Types:     in.Types,
		Duration:  in.Duration,
		Area:      in.Area,
	}

	if err := db.InsertVideo(l.ctx, &dbvideo); err != nil {
		return nil, err
	}

	// 通过rocketmq通知到search服务
	str, err := json.Marshal(dbvideo)
	if err != nil {
		return nil, err
	}
	msg := &primitive.Message{
		Topic: constant.ROCKETMQ_ADDVIDEO_TOPIC,
		Body:  str,
	}
	_, err = l.svcCtx.MQProducer.SendSync(l.ctx, msg)
	if err != nil {
		logx.Error("发送rocketmq消息失败")
		return nil, err
	}
	return &video.AddVideoResponse{}, nil
}
