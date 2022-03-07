package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const dsn = "host=localhost user=phobyjun password=password dbname=goropbox port=5432 TimeZone=Asia/Seoul"

func Init() *gorm.DB {
	session := getSession()

	return session
}

func getSession() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
