package controllers

import (
	"github.com/astaxie/beego"
)


//type BaseController struct {
//	beego.Controller
//	IsLogin bool
//	User   models.User
//}

type MainController struct {
	beego.Controller
}

type SubController struct {
	beego.Controller
}

//func (c *BaseController) Prepare(){
//	c.IsLogin = false
//	key := c.GetSession("SESSION_KEY")
//	if key != nil {
//		if u, ok := key.(models.User); ok {
//			c.User = u
//			c.IsLogin = true
//		}
//	}
//}

func (c *MainController) Get() {
	c.Data["Website"] = "nibaba"
	c.Data["Email"] = "lufeng@gmail.com"
	c.Data["Username"] = "lufeng"
	c.TplName = "index.html"
}

// 执行子域名扫描任务和展示
func (c *SubController) get() {

}

