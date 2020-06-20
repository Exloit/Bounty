package routers

import (
	"Bounty/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
    // login and register remove
	//beego.Router("/login", &controllers.LoginController{})
    //beego.Router("/register", &controllers.RegisterControllers{})
    beego.Router("/subdomains", &controllers.SubController{})
    beego.Router("/masscan", &controllers.MasscanController{})
}
