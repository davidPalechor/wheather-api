package routers

import (
	"WheatherAPI/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/get", &controllers.ApiController{})
}