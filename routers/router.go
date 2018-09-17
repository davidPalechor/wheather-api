// @APIVersion 1.0.0
// @Title Weather API
// @Description API for viewing weather actual status.
// @Contact sergio.palechor@globant.com

package routers

import (
	"WheatherAPI/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/weather",
				beego.NSInclude(&controllers.WeatherController{}),
			),
			beego.NSNamespace("/scheduler/weather",
				beego.NSInclude(&controllers.ScheduleController{}),
			),
		)
	beego.AddNamespace(ns)
}