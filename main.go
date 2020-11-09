package main

import (
	"Bounty/controllers"
	"Bounty/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

const Dsn = "root:root123456@tcp(127.0.0.1:3306)/bounty?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	//defer db
	if err := model.Db.AutoMigrate(&model.Users{}); err != nil {
		CheckErr(err)
	}
	router := gin.Default()
	//router.GET("/user", controllers.UserQuery)
	router.POST("/register", controllers.UserRegister)
	router.POST("/login", controllers.UserLogin)
	router.DELETE("/user", controllers.UserDelete)
	router.Run(":9999")
}

func CheckErr(err error) {
	if err != nil {
		fmt.Printf(err.Error())
	}
}
