package db

import (
	"context"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserId      uint64 `gorm:"column:user_id;type:bigint;not null"`
	Url         string `gorm:"column:url;type:varchar(500); not null"`
	Thumbnail   string
	Title       string
	Types       int32
	Duration    string
	Area        string
	description string
}

func (Video) TableName() string {
	return "t_video"
}

func InsertVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

func UpdateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Updates(video).Error
}

func QueryVideoByVideoId(ctx context.Context, videoId uint64) (*Video, error) {
	res := &Video{}
	if err := DB.WithContext(ctx).Where("id = ?", videoId).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
