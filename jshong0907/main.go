package main

import (
	"fmt"

	"github.com/jshong0907/GoropBox/jshong0907/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=127.0.0.1 user=hong dbname=GoropBox"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// Create
	db.Create(&models.User{Email: "jshong0907@gmail.com", NickName: "이름"})

	// Read
	var user models.User
	db.First(&user, 1) // find product with integer primary key
	fmt.Println(user.NickName)
	db.First(&user, "Email = ?", "jshong0907@gmail.com") // find product with code D42

	// Update - update product's price to 200
	db.Model(&user).Update("NickName", "홍준식")
	// Update - update multiple fields

	// Delete - delete product
	db.Where("email = ?", "jshong0907@gmail.com").Delete(&user)
}
