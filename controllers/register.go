package controllers

import (
	"Bounty/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type RegisterControllers struct {
	beego.Controller
}

func (c *RegisterControllers) Get(){
	c.TplName = "register.html"
}

func (c *RegisterControllers) Post(){
	username := strings.TrimSpace(c.GetString("username"))
	password := strings.TrimSpace(c.GetString("password"))

	o := orm.NewOrm()
	var user models.User
	if username == "" || password == ""{

	}else{
		user.Name = username
		user.Password = password
		_, err := o.Insert(&user)
		if err != nil {
			c.Data["json"] = &JSONStruct{200, true ,"register error"}
			c.ServeJSON()
		}
		c.Data["json"] = &JSONStruct{200, true,"register success"}
		c.ServeJSON()
	}
}
