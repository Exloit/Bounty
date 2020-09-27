package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const Dsn = "root:root123456@tcp(127.0.0.1:3306)/bounty?charset=utf8mb4&parseTime=True&loc=Local"

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Printf(err.Error())
	}
}
