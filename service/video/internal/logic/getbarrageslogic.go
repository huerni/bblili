package logic

import (
	"bblili/service/video/internal/common/constant"
	"bblili/service/video/internal/db"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBarragesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBarragesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBarragesLogic {
	return &GetBarragesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBarragesLogic) GetBarrages(in *video.GetBarragesRequest) (*video.GetBarragesResponse, error) {

	barrageInfos := make([]*video.BarrageInfo, 0)
	key := fmt.Sprintf(constant.BARRAGE_SORTEDSET, strconv.FormatUint(in.VideoId, 10))
	isexist, err := l.svcCtx.RedisClient.Exists(key)
	if err != nil {
		return nil, err
	}

	if isexist {
		// 加分布式锁
		mutex := l.svcCtx.RedisSync.NewMutex(constant.LOCK_BARRAGE_SORTEDEST)
		// 对key进行
		if err := mutex.Lock(); err != nil {
			panic(err)
		}
		defer mutex.Unlock()
		key := fmt.Sprintf(constant.BARRAGE_SORTEDSET, strconv.FormatUint(in.VideoId, 10))
		isexist, err := l.svcCtx.RedisClient.Exists(key)
		if err != nil {
			return nil, err
		}
		if isexist {
			redisres, err := l.svcCtx.RedisClient.ZrangebyscoreWithScores(key, int64(in.CurrentTime), int64(in.CurrentTime+60))
			if err != nil {
				return nil, err
			}
			// 拼接成结果
			for _, redisval := range redisres {
				var barrage db.Barrage
				err = json.Unmarshal([]byte(redisval.Key), &barrage)
				if err != nil {
					return nil, err
				}
				barrageInfos = append(barrageInfos, &video.BarrageInfo{
					UserId:      barrage.UserId,
					VideoId:     barrage.VideoId,
					Content:     barrage.Content,
					BarrageTime: barrage.BarrageTime,
				})
			}
		}
	} else {
		barrages, err := db.QueryBarrageByVideoIdAndCurrentTime(l.ctx, in.VideoId, in.CurrentTime)
		if err != nil {
			return nil, err
		}

		// todo 用lua脚本将所有barrage插入redis中
		mutex := l.svcCtx.RedisSync.NewMutex(constant.LOCK_BARRAGE_SORTEDEST)
		// 对key进行
		if err := mutex.Lock(); err != nil {
			panic(err)
		}
		defer mutex.Unlock()
		for _, barrage := range barrages {
			bytes, err := json.Marshal(barrage)
			if err != nil {
				return nil, err
			}
			isadd, err := l.svcCtx.RedisClient.Zadd(key, int64(barrage.BarrageTime), string(bytes))
			if isadd == false || err != nil {
				logx.Errorf("redis插入数据失败")
			}
			barrageInfos = append(barrageInfos, &video.BarrageInfo{
				UserId:      barrage.UserId,
				VideoId:     barrage.VideoId,
				Content:     barrage.Content,
				BarrageTime: barrage.BarrageTime,
			})
		}
	}

	return &video.GetBarragesResponse{BarrageInfos: barrageInfos}, nil
}
