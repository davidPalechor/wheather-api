package routers

import (
	"WheatherAPI/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/weather", &controllers.WeatherController{})
	beego.Router("/scheduler/weather", &controllers.ScheduleController{})
}