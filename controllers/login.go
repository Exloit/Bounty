package controllers

import (
	"Bounty/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type JSONStruct struct{
	Code int
	Success bool
	Msg string
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get(){
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	var user models.User
	name := strings.TrimSpace(c.GetString("username"))
	password := strings.TrimSpace(c.GetString("password"))
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("Name", name).Filter("Password", password).One(&user)
	if err != nil {
		fmt.Println("登录出错")
		c.Data["json"] = JSONStruct{200, true,"login failed"}
		c.ServeJSON()
	}else {
		fmt.Println(name)
		c.SetSession("sid", user)
		c.Redirect("/admin", 200)
	}
}