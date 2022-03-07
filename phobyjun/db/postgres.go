package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"phobyjun/model"
)

const dsn = "host=localhost user=phobyjun password=password dbname=goropbox port=5432 TimeZone=Asia/Seoul"

var db *gorm.DB

func Init() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(
		&model.User{},
		&model.File{},
	)
}

func DBManager() *gorm.DB {
	return db
}
