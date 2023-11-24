package db

import (
	"context"
	"gorm.io/gorm"
)

var VideoCollectionTable = "t_video_collection"

type VideoCollection struct {
	gorm.Model
	VideoId uint64
	UserId  uint64
	GroupId uint64
}

func (VideoCollection) TableName() string {
	return "t_video_collection"
}

func InsertVideoCollection(ctx context.Context, collection *VideoCollection) error {
	return DB.WithContext(ctx).Create(collection).Error
}

func UpdateVideoCollection(ctx context.Context, collection *VideoCollection) error {
	return DB.WithContext(ctx).Updates(collection).Error
}

func DeleteVideoCollection(ctx context.Context, ID uint64) error {
	res := make([]*VideoCollection, 0)
	return DB.WithContext(ctx).Delete(&res, ID).Error
}

func QueryVideoCollectionByVideoIdAndUserIdAndGroupId(ctx context.Context, videoId uint64, userId uint64, groupId uint64) ([]*VideoCollection, error) {
	res := make([]*VideoCollection, 0)
	if err := DB.WithContext(ctx).Where("video_id = ? and user_id = ? and group_id = ?", videoId, userId, groupId).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func CountVideoCollectionByVideoId(ctx context.Context, videoId uint64) (*int64, error) {
	var count int64
	if err := DB.Table("t_video_collection").Where("video_id = ?", videoId).Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}

func QueryVideoCollectionByVideoIdAdnUserId(ctx context.Context, videoId uint64, userId uint64) ([]*VideoCollection, error) {
	res := make([]*VideoCollection, 0)
	if err := DB.WithContext(ctx).Where("video_id = ? and user_id = ?", videoId, userId).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
