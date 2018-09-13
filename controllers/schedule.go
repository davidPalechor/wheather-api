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
	err := tasks.WeatherTask(city, country)
	if err != nil {
		fmt.Println(err)
	}
	// Response to client
	return
}
