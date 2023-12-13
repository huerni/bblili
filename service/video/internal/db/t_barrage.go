package db

import (
	"context"
	"gorm.io/gorm"
)

type Barrage struct {
	gorm.Model
	UserId      uint64
	VideoId     uint64
	Content     string
	BarrageTime int32
}

func (Barrage) TableName() string {
	return "t_barrage"
}

func InsertDanMu(ctx context.Context, danMu *Barrage) error {
	return DB.WithContext(ctx).Create(danMu).Error
}

func QueryBarrageByVideoId(ctx context.Context, videoId uint64) ([]*Barrage, error) {
	res := make([]*Barrage, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", videoId).Order("barrage_time ASC").Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func QueryBarrageByVideoIdAndCurrentTime(ctx context.Context, videoId uint64, currentTime int32) ([]*Barrage, error) {
	res := make([]*Barrage, 0)
	if err := DB.WithContext(ctx).Where("video_id = ? and barrage_time >= ? and barrage_time < ?", videoId, currentTime, currentTime+60).Order("barrage_time ASC").Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
