// @Date  : 2020/6/17
// @Author: Lu Feng
// @File  : Masscan.go
// @Description: 

package models

import "time"

type Masscan struct {
    Ip   string
    CreateTime time.Time
    FinishedTime time.Time
    DeleteTime time.Time
}
