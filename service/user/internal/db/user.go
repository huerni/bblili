package db

import (
	"context"
	"gorm.io/gorm"
)

var userTableName = "t_user"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(254);not null"`
	Password string `gorm:"column:password;type:varchar(254);not null"`
	Phone    string `gorm:"column:phone;type:varchar(100)"`
	Email    string `gorm:"column:email;type:varchar(100)"`
}

func (u *User) TableName() string {
	return userTableName
}

func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUserByPhone(ctx context.Context, phone string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("phone = ?", phone).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUserByEmail(ctx context.Context, email string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("email = ?", email).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUserById(ctx context.Context, ID uint64) (*User, error) {
	var res User
	if err := DB.WithContext(ctx).Where("id = ?", ID).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func InsertUser(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(user).Error
}

func UpdateUser(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Updates(user).Error
}
