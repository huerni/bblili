package logic

import (
	"bblili/service/video/internal/common/constant"
	"bblili/service/video/internal/db"
	"bblili/service/video/internal/svc"
	"bblili/service/video/video"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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

	videocomment := db.VideoComment{
		Model:       gorm.Model{},
		VideoId:     in.VideoId,
		UserId:      in.UserId,
		Comment:     in.Comment,
		ReplyUserId: in.ReplyUserId,
		RootId:      in.RootId,
	}

	if err := db.InsertVideoComment(l.ctx, &videocomment); err != nil {
		return nil, err
	}

	go func() {
		// 加分布式锁
		lockkey := fmt.Sprintf(constant.LOCK_VIDEOCOMMENT_SORTEDSET, in.VideoId, in.RootId)
		mutex := l.svcCtx.RedisSync.NewMutex(lockkey)
		// 对key进行
		if err := mutex.Lock(); err != nil {
			logx.Errorf("获得redis分布式锁失败-%s", err)
		}
		defer func() {
			// 释放互斥锁
			if ok, err := mutex.Unlock(); !ok || err != nil {
				logx.Errorf("释放redis分布式锁失败-%s", err)
			}
		}()
		key := fmt.Sprintf(constant.VIDEOCOMMENT_SORTEDSET, in.VideoId, in.RootId)
		bytes, err := json.Marshal(videocomment)
		if err != nil {
			logx.Errorf("redis插入数据失败-%s", key)
		}

		isadd, err := l.svcCtx.RedisClient.Zadd(key, videocomment.CreatedAt.Unix(), string(bytes))
		if isadd == false || err != nil {
			logx.Errorf("redis插入数据失败-%s-%s", key, err)
		}
	}()

	return &video.AddVideoCommentResponse{}, nil
}
