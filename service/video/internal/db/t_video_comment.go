package db

import (
	"context"
	"gorm.io/gorm"
)

type VideoComment struct {
	gorm.Model
	VideoId     uint64
	UserId      uint64
	Comment     string
	ReplyUserId uint64
	RootId      uint64
}

func (VideoComment) TableName() string {
	return "t_video_comment"
}

func InsertVideoComment(ctx context.Context, comment *VideoComment) error {
	return DB.WithContext(ctx).Create(comment).Error
}

func UpdateVideoComment(ctx context.Context, comment *VideoComment) error {
	return DB.WithContext(ctx).Updates(comment).Error
}

func PageQueryVideoCommentByVideoId(ctx context.Context, page int32, size int32, videoId uint64) ([]*VideoComment, error) {
	res := make([]*VideoComment, 0)
	var sqlstr string

	sqlstr = "select * from t_video_comment where id in (select id from t_video_comment where video_id = ? and root_id = 0 order by created_at desc) limit ? offset ?"
	if err := DB.WithContext(ctx).Raw(sqlstr, videoId, size, (page-1)*size).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func QueryVideoChildCommentByVideoId(ctx context.Context, videoId uint64, rootId uint64) ([]*VideoComment, error) {
	res := make([]*VideoComment, 0)

	if err := DB.WithContext(ctx).Where("video_id = ? and root_id = ?", videoId, rootId).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
