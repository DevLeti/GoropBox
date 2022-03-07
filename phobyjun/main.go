package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"phobyjun/db"
	"phobyjun/model"
)

func main() {
	db.Init()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello Echo!")
	})
	e.GET("/api/v2/users", getUsers)
	e.POST("/api/v2/users", createUsers)
	e.Logger.Fatal(e.Start(":8080"))
}

func getUsers(c echo.Context) error {
	tx := db.Session
	var users []model.User
	tx.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func createUsers(c echo.Context) error {
	tx := db.Session
	user := model.User{
		Username: "junseok",
		Password: "0000",
		Email:    "phobyjun@gmail.com",
	}
	tx.Create(&user)
	return c.JSON(http.StatusOK, user)
}
