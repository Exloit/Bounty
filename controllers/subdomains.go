package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
)


type SubController struct {
	beego.Controller
}

func (c *SubController)  Get(){
	c.TplName = "subdomains.html"
}

func (c *SubController) Post(){
	domain := c.GetString("domains")
	resp := &JSONStruct{200, "test "+ domain}
	go StartSubDomainsScan(domain)
	c.Data["json"] = &resp
	c.ServeJSON()
}

func StartSubDomainsScan(target string){
	time.Sleep(1000)
	fmt.Println(target)
}
