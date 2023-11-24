package db

import (
	"context"
	"gorm.io/gorm"
)

var VideoCoinTableName = "t_video_coin"

type VideoCoin struct {
	gorm.Model
	UserId  uint64
	VideoId uint64
	Amount  int
}

func (v *VideoCoin) TableName() string {
	return VideoCoinTableName
}

func InsertVideoCoin(ctx context.Context, coin *VideoCoin) error {
	return DB.WithContext(ctx).Create(coin).Error
}

func UpdateVideoCoin(ctx context.Context, coin *VideoCoin) error {
	return DB.WithContext(ctx).Updates(coin).Error
}

func QueryVideoCoinByVideoIdAndUserId(ctx context.Context, videoId uint64, userID uint64) ([]*VideoCoin, error) {
	res := make([]*VideoCoin, 0)
	if err := DB.WithContext(ctx).Where("video_id = ? and user_id = ?", videoId, userID).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func CountVideoCoinByVideoId(ctx context.Context, videoId uint64) (*int64, error) {
	var res int64
	if err := DB.Table("t_video_coin").Where("video_id = ?", videoId).Count(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
