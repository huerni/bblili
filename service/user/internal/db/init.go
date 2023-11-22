package db

import "bblili/service/user/internal/config"
import "gorm.io/gorm"
import "gorm.io/driver/mysql"

var DB *gorm.DB

func Init(c *config.Config) error {
	var err error
	DB, err = gorm.Open(mysql.Open(c.DSN))
	if err != nil {
		return err
	}

	return nil
}
