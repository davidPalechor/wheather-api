package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"WheatherAPI/tasks"
)


type ScheduleController struct {
	beego.Controller
}


func (c *ScheduleController) Put() {
	city := c.GetString("city")
	country := c.GetString("country")
	fmt.Print(city, country)
	tasks.WeatherTask(city, country)

	c.Ctx.ResponseWriter.WriteHeader(200)
	// Response to client
	return
}
