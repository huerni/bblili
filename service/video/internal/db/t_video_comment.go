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

	sqlstr = "select id, create_at, update_at, delete_at, video_id, user_id, comment, reply_user_id, root_id from t_video_comment where video_id = ? and root_id is null and id = (select id from t_video_comment limit ? offset ? order by create_at)"
	if err := DB.WithContext(ctx).Raw(sqlstr, videoId, size, (page-1)*size).Scan(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func QueryVideoChildCommentByVideoId(ctx context.Context, videoId uint64, rootId uint64) ([]*VideoComment, error) {
	res := make([]*VideoComment, 0)

	if err := DB.WithContext(ctx).Where("video_id = ? and root_id = ?", videoId, rootId).Error; err != nil {
		return nil, err
	}

	return res, nil
}
