package main

import (
	"fmt"
	"github.com/devleti/goropbox/devleti/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "host=localhost user=leti password=password dbname=goropbox port=5432 TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect DB")
	}

	var migrateAnswer string
	fmt.Println("Need migrate? Y or N : ")

	fmt.Scanln(&migrateAnswer)
	switch migrateAnswer {
	case "Y", "y":
		checkMigrate(db)
	default:
		fmt.Println("Skip migrate.")
	}
	//err = db.AutoMigrate(
	//	&model.User{},
	//	&model.File{},
	//)
	//if err != nil {
	//	panic("Failed to migrate")
	//}
}

func checkMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.File{},
	)
	if err != nil {
		panic("Failed to migrate")
	}
}
