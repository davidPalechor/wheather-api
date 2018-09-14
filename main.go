package main

import (
	_ "WheatherAPI/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/toolbox"
)

func init() {
	orm.Debug = true
    orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:root@tcp(127.0.0.1:3306)/weather?charset=utf8",
	)
}

func main() {
	defer toolbox.StopTask()
	beego.Run()
}
