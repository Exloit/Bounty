package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"math/rand"
	"github.com/OWASP/Amass/v3/config"
	"github.com/OWASP/Amass/v3/systems"
	"github.com/OWASP/Amass/v3/datasrcs"
	"github.com/OWASP/Amass/v3/enum"
)


type SubController struct {
	beego.Controller
}

func (c *SubController)  Get(){
	c.TplName = "subdomains.html"
}

func (c *SubController) Post(){
	domain := c.GetString("domains")
	resp := &JSONStruct{200, true,"test "+ domain}
	go StartSubDomainsScan(domain)
	c.Data["json"] = &resp
	c.ServeJSON()
}

func StartSubDomainsScan(target string){

	rand.Seed(time.Now().UTC().UnixNano())

	cfg := config.NewConfig()
	cfg.AddDomain(target)

	sys, err := systems.NewLocalSystem(cfg)
	if err != nil{
		return
	}
	sys.SetDataSources(datasrcs.GetAllSources(sys))

	e := enum.NewEnumeration(cfg, sys)

	if e == nil {
		return
	}

	defer e.Close()

	e.Start()
	for _, o := range e.ExtractOutput(nil){
		fmt.Println(o.Name)
	}

}


