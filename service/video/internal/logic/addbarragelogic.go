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
	"strconv"
)

type AddBarrageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBarrageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBarrageLogic {
	return &AddBarrageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加弹幕
func (l *AddBarrageLogic) AddBarrage(in *video.AddBarrageRequset) (*video.AddBarrageResponse, error) {

	barrage := db.Barrage{
		Model:       gorm.Model{},
		UserId:      in.BarrageInfo.UserId,
		VideoId:     in.BarrageInfo.VideoId,
		Content:     in.BarrageInfo.Content,
		BarrageTime: in.BarrageInfo.BarrageTime,
	}

	if err := db.InsertDanMu(l.ctx, &barrage); err != nil {
		return nil, err
	}

	go func() {
		// 加分布式锁
		mutex := l.svcCtx.RedisSync.NewMutex(constant.LOCK_BARRAGE_SORTEDEST)
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

		key := fmt.Sprintf(constant.BARRAGE_SORTEDSET, strconv.FormatUint(in.BarrageInfo.VideoId, 10))
		bytes, err := json.Marshal(barrage)
		if err != nil {
			logx.Errorf("redis插入数据失败-%s", key)
		}

		isadd, err := l.svcCtx.RedisClient.Zadd(key, int64(barrage.BarrageTime), string(bytes))
		if isadd == false || err != nil {
			logx.Errorf("redis插入数据失败-%s-%s", key, err)
		}
	}()

	return &video.AddBarrageResponse{}, nil
}
