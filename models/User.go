package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//	TODO: 这里后面需要改成配置文件读取
func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root123456@tcp(localhost:3306)/bounty?charset=utf8mb4")
	orm.RegisterModel(new(User))

	orm.RunSyncdb("default", false, true)
}

type User struct {
	Id int
	JobId int
	Name string `orm:"unique"`
	Password string
	DepartName string
	Permission  string
}
