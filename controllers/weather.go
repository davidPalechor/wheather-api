package controllers

import (
	"github.com/astaxie/beego"
	"WheatherAPI/utils"
)


type WeatherController struct {
	beego.Controller
}

// @Title Get
// @Summary Get
// @Description Retrieve weather info about a City/Country
// @Param   city	query	string	true	"Name of the city capitilized e.g: Bogota"
// @Param   country	query	string	true	"Code of country lowercased e.g: co (Colombia), uk (United Kingdom), ca (Canada)"
// @Success 200 {string} "Success"
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router / [get]
func (c *WeatherController) Get() {
	city := c.GetString("city")
	country := c.GetString("country")
	response,_ := utils.WeatherReporter(city, country)

	// Response to client
	c.Data["json"] = response
	c.ServeJSON()
}
