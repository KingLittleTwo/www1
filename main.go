package main

import (
	_ "ww1/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/ww_ago?charset=utf8", 30)
	beego.Run()
}

