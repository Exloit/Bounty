// @Date  : 2020/6/17
// @Author: Lu Feng
// @File  : Masscan.go
// @Description: 

package controllers

import (
    "github.com/astaxie/beego"
    //"encoding/json"
    "fmt"
    "github.com/dean2021/go-masscan"
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
    //var ip Ip
    ip := c.GetString("ip")
    fmt.Println(ip)
    //err := json.Unmarshal(data, &ip)
    //if err != nil {
    //    fmt.Println("json.Unmarshal is err:", err.Error())
    //}
    go StartMasscan(ip)
    resp := &JSONStruct{200 , true, "get ip" }
    c.Data["json"] = &resp
    c.ServeJSON()
}

// 删除一个masscan扫描
func (c *MasscanController) Delete(){

}

func StartMasscan(ip string){
    fmt.Println(ip)
    m := masscan.New()

    m.SetPorts("0-4000")

    m.SetRanges(ip)

    m.SetRate("300")

    err := m.Run()

    if err != nil {
        fmt.Println("scanner failed", err)
        return
    }

    results, err := m.Parse()
    if err != nil {
        fmt.Println("Parse scanner result:", err)
        return
    }

    for _, result := range results {
        fmt.Println(result)
    }

}