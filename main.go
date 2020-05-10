package main

import (
	"Bounty/models"
	_ "Bounty/routers"
	"encoding/gob"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "session-id"
	beego.Run()
}

