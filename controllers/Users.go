package controllers

import (
	"Bounty/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {

	var newUser model.Users
	if err := c.ShouldBindJSON(&newUser); err != err {
		fmt.Printf(err.Error())
	}
	model.Db.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func UserLogin(c *gin.Context) {
	var userInfo model.Users
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		fmt.Printf(err.Error())
	}
	model.Db.Find(&userInfo, "user_name = ? AND password = ?", userInfo.UserName, userInfo.Password)
	if userInfo.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"status": "Failed"})
	}
}

func UserUpdate(c *gin.Context) {
	var userInfo model.Users
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		fmt.Printf(err.Error())
		c.JSON(http.StatusForbidden, gin.H{"status": "failed"})
	}
	model.Db.Save(&userInfo)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func UserDelete(c *gin.Context) {

}
