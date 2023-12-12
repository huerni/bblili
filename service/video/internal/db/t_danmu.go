package db

import (
	"context"
	"gorm.io/gorm"
)

type DanMu struct {
	gorm.Model
	UserId    uint64
	VideoId   uint64
	Content   string
	DanmuTime string
}

func (DanMu) TableName() string {
	return "t_danmu"
}

func InsertDanMu(ctx context.Context, danMu *DanMu) error {
	return DB.WithContext(ctx).Create(danMu).Error
}

func QueryDanMuByVideoId(ctx context.Context, videoId uint64) ([]*DanMu, error) {
	res := make([]*DanMu, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", videoId).Order("danmu_time ASC").Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
