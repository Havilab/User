package database

import (
	"github.com/ekart/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UserDb *gorm.DB

func DbConnection(dsn string) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database not connected")
	}

	db.AutoMigrate(model.User{}, model.Orders{})
	UserDb = db
}
func Db() *gorm.DB {
	return UserDb
}
