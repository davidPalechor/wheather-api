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
			"%s:%s@tcp(%s)/%s?charset=utf8",
			beego.AppConfig.String("mysqluser"),
			beego.AppConfig.String("mysqlpswd"),
			os.Getenv("DB_HOST"),
			beego.AppConfig.String("mysqldb"),
		),
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
