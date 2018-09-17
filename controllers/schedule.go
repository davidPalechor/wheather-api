package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"WheatherAPI/tasks"
)


type ScheduleController struct {
	beego.Controller
}

// @Title Put
// @Summary Put
// @Description Retrieve weather info about a City/Country
// @Param   city	body	string	true	"Name of the city capitilized e.g: Bogota"
// @Param   country	body	string	true	"Code of country lowercased e.g: co (Colombia), uk (United Kingdom), ca (Canada)"
// @Success 200 {string} "Success"
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router / [put]
func (c *ScheduleController) Put() {
	city := c.GetString("city")
	country := c.GetString("country")
	fmt.Print(city, country)
	tasks.WeatherTask(city, country)

	c.Ctx.ResponseWriter.WriteHeader(200)
	// Response to client
	return
}
