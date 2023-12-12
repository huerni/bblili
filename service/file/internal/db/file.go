package db

import (
	"context"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Url      string
	Type     string
	Md5      string
	UUID     string
	UploadID string
	Uploaded bool
}

func (f *File) TableName() string {
	return "t_file"
}

func QueryFileByMD5(ctx context.Context, md5 string) (*File, error) {
	res := &File{}
	if err := DB.WithContext(ctx).Where("md5 = ?", md5).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func InsertFile(ctx context.Context, file *File) error {
	return DB.WithContext(ctx).Create(file).Error
}

func UpdateFile(ctx context.Context, file *File) error {
	return DB.WithContext(ctx).Updates(file).Error
}
