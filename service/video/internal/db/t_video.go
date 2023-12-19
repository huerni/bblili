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
	Duration    int32
	Area        int32
	Description string
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

func QueryVideo(ctx context.Context, page int32, size int32) ([]*Video, error) {
	res := make([]*Video, 0)

	sqlstr := "select * from t_video where id = (select id from t_video order by created_at limit ? offset ?)"
	if err := DB.WithContext(ctx).Raw(sqlstr, page, (page-1)*size).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func QueryVideoByUserId(ctx context.Context, page int32, size int32, userId uint64) ([]*Video, error) {
	res := make([]*Video, 0)

	sqlstr := "select * from t_video where user_id = ? and  id = (select id from t_video order by created_at limit ? offset ?)"
	if err := DB.WithContext(ctx).Raw(sqlstr, userId, page, (page-1)*size).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func QueryVideoByArea(ctx context.Context, page int32, size int32, area int32) ([]*Video, error) {
	res := make([]*Video, 0)

	sqlstr := "select * from t_video where area = ? and  id = (select id from t_video order by created_at limit ? offset ?)"
	if err := DB.WithContext(ctx).Raw(sqlstr, area, page, (page-1)*size).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
