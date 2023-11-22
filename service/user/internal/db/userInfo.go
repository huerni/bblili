package db

import (
	"context"
	"gorm.io/gorm"
)

var userInfoTableName = "t_user_info"

type UserInfo struct {
	gorm.Model
	Nick   string `gorm:"column:nick;type:varchar(100);not null"`
	Avatar string `gorm:"column:avatar;type:varchar(1024)"`
	Sign   string `gorm:"column:sign;type:text"`
	Gender int32  `gorm:"column:gender;type:integer"`
	Birth  string `gorm:"column:birth;type:varchar(20)"`
}

func QueryByUserId(ctx context.Context, userId uint64) ([]*UserInfo, error) {
	res := make([]*UserInfo, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func InsertUserInfo(ctx context.Context, userInfo *UserInfo) error {
	return DB.WithContext(ctx).Create(userInfo).Error
}

func UpdateUserInfo(ctx context.Context, userInfo *UserInfo) error {
	return DB.WithContext(ctx).Updates(userInfo).Error
}
