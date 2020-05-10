package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/bounty?charset=utf8mb4")
	orm.RegisterModel(new(User))

	orm.RunSyncdb("default", false, true)
}

type User struct {
	Id int
	Name string `orm:"unique"`
	Password string
}
