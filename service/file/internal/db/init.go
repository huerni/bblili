package db

import (
	"bblili/service/file/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(c *config.Config) error {
	var err error
	DB, err = gorm.Open(mysql.Open(c.DSN))
	if err != nil {
		return err
	}

	return nil
}
