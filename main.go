package main

import (
	_ "WheatherAPI/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/toolbox"
	"os"
	"fmt"
)

func init() {
	if beego.BConfig.RunMode == "dev"{
		orm.Debug = true
	}
    orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(
		"default",
		"mysql",
		fmt.Sprintf(
			"root:root@tcp(%s)/weather?charset=utf8",
			os.Getenv("DB_HOST"),
		),
	)
}

func main() {
	defer toolbox.StopTask()
	fmt.Printf(
		"root:root@tcp(%s)/weather?charset=utf8",
		os.Getenv("DB_HOST"),
	)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
