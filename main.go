package main

import (
	_ "WheatherAPI/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/toolbox"
)

func init() {
	if beego.BConfig.RunMode == "dev"{
		orm.Debug = true
	}
    orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:root@tcp(127.0.0.1:3306)/weather?charset=utf8",
	)
}

func main() {
	defer toolbox.StopTask()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
