package controllers

import (
	"github.com/astaxie/beego"
	"WheatherAPI/utils"
)


type WeatherController struct {
	beego.Controller
}


func (c *WeatherController) Get() {
	city := c.GetString("city")
	country := c.GetString("country")
	response := utils.WeatherReporter(city, country)

	// Response to client
	c.Data["json"] = response
	c.ServeJSON()
}
