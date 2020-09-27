package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	DepartId    int32
	DepartName  string
	Email       string
	Permissions string
	RealName    string
	RoleName    string
	UserName    string `gorm:"unique"`
	Password    string
	JobId       string
	Avatar      string
}
