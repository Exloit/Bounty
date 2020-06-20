// @Date  : 2020/6/17
// @Author: Lu Feng
// @File  : Masscan.go
// @Description: 

package controllers

import (
    "github.com/astaxie/beego"
    "encoding/json"
    "fmt"
)

type MasscanController struct {
    beego.Controller
}

type Ip struct {
    Ip string
}

// 获取所有masscan扫描记录
func (c *MasscanController) Get(){
    c.TplName = "masscan.html"
}

// 新增一个masscan扫描
func (c *MasscanController) Post(){
    var ip Ip
    data := c.Ctx.Input.RequestBody
    err := json.Unmarshal(data, &ip)
    if err != nil {
        fmt.Println("json.Unmarshal is err:", err.Error())
    }
    resp := &JSONStruct{200 , true, "get ip" }
    c.Data["json"] = &resp
    c.ServeJSON()
}

// 删除一个masscan扫描
func (c *MasscanController) Delete(){

}