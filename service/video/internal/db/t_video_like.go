package db

import (
	"context"
	"gorm.io/gorm"
)

type VideoLike struct {
	gorm.Model
	UserId  uint64
	VideoId uint64
}

func (VideoLike) TableName() string {
	return "t_video_like"
}

func InsertVideoLike(ctx context.Context, like *VideoLike) error {
	return DB.WithContext(ctx).Create(like).Error
}

func UpdateVideoLike(ctx context.Context, like *VideoLike) error {
	return DB.WithContext(ctx).Updates(like).Error
}

func DeleteVideoLike(ctx context.Context, ID uint64) error {
	res := make([]*VideoLike, 0)
	return DB.WithContext(ctx).Delete(&res, ID).Error
}

func QueryVideoLikeByUserIdAndVideoId(ctx context.Context, userId uint64, videoId uint64) ([]*VideoLike, error) {
	res := make([]*VideoLike, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", userId, videoId).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func CountVideoLikeByVideoId(ctx context.Context, videoId uint64) (*int64, error) {
	var count int64
	if err := DB.Table("t_video_like").Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}
