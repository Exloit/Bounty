package routers

import (
	"Bounty/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var LoginCheck = func(ctx *context.Context) {
	userName := ctx.Input.Session("sid")
	fmt.Println(userName)
	if userName == nil {
		ctx.Redirect(302, "/login")
	}
}

func init() {
	beego.InsertFilter("/admin*", beego.BeforeRouter, LoginCheck)
    beego.Router("/admin", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
    beego.Router("/register", &controllers.RegisterControllers{})
    beego.Router("/admin/subdomain", &controllers.SubController{})
}
