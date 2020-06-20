package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"os/exec"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/log"
	"math/rand"
	"github.com/OWASP/Amass/config"
	"github.com/OWASP/Amass/systems"
	"github.com/OWASP/Amass/datasrcs"
	"github.com/OWASP/Amass/enum"
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
	time.Sleep(1000)
	//args := "-d " + target
	//cmd := exec.Command("amass", args)
	//err := cmd.Start()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = cmd.Wait()

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


