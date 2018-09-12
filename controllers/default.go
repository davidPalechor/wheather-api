package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"time"
)


type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


type ApiController struct {
	beego.Controller
}


func (c *ApiController) Get() {
	var jsonResponse map[string]interface{}

	city := c.GetString("city")
	country := c.GetString("country")

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=1508a9a4840a5574c822d70ca2132032", city, country)
	req := httplib.Get(url)
	data, _ := req.String()

	err := json.Unmarshal([]byte(data), &jsonResponse)

	if err != nil{
		c.Data["json"] = `{"error": 0}`
	}

	main := jsonResponse["main"].(map[string] interface {})
	sys := jsonResponse["sys"].(map[string] interface{})
	wind := jsonResponse["wind"].(map[string] interface{})
	coord := jsonResponse["coord"].(map[string] interface{})
	sunrise := time.Unix(int64(sys["sunrise"].(float64)), 0)
	sunset := time.Unix(int64(sys["sunset"].(float64)), 0)
	requestedTime := time.Unix(int64(jsonResponse["dt"].(float64)), 0)

	c.Data["json"] = map[string]interface{}{
		"location_name": fmt.Sprintf(
			"%s, %s",
			jsonResponse["name"],
			sys["country"],
		),
		"temperature": fmt.Sprintf(
			"%.2f °C",
			main["temp"].(float64) - 273.15,
		),
		"wind": fmt.Sprintf(
			"%s, %.2f m/s, %s",
			beaufortScale(wind["speed"].(float64)),
			wind["speed"],
			windDirection(wind["deg"].(float64)),
		),
		"pressure": fmt.Sprintf(
			"%.0f hPa",
			main["pressure"],
		),
		"humidity": fmt.Sprintf(
			"%.0f%%",
			main["humidity"],
		),
		"geo_coordinates": []float64{
			coord["lat"].(float64),
			coord["lon"].(float64),
		},
		"sunrise": fmt.Sprintf(
			"%02d:%02d",
			sunrise.Hour(),
			sunrise.Minute(),
		),
		"sunset": fmt.Sprintf(
			"%02d:%02d",
			sunset.Hour(),
			sunset.Minute(),
		),
		"requested_time": fmt.Sprintf(
			"%d-%02d-%02d %02d:%02d:%02d",
			requestedTime.Year(),
			requestedTime.Month(),
			requestedTime.Day(),
			requestedTime.Hour(),
			requestedTime.Minute(),
			requestedTime.Second(),
		),
	}
	c.ServeJSON()
}


func beaufortScale(speed float64)(scale string) {
	if speed < 0.3 {
		scale = "Calm"
	} else if 0.3 <= speed && speed <= 1.5 {
		scale = "Light Air"
	} else if 1.6 <= speed && speed <= 3.3 {
		scale = "Light Breeze"
	} else if 3.4 <= speed && speed <= 5.5 {
		scale = "Gentle Breeze"
	} else if 5.6 <= speed && speed <= 7.9 {
		scale = "Moderate Breeze"
	} else if 8.0 <= speed && speed <= 10.7 {
		scale = "Fresh Breeze"
	} else if 10.8 <= speed && speed <= 13.8 {
		scale = "Strong Breeze"
	} else if 13.9 <= speed && speed <= 17.1 {
		scale = "High Wind"
	} else if 17.2 <= speed && speed <= 20.7 {
		scale = "Fresh Gale"
	} else if 20.8 <= speed && speed <= 24.4 {
		scale = "Strong Gale"
	} else if 24.5 <= speed && speed <= 28.4 {
		scale = "Storm"
	} else if 28.5 <= speed && speed <= 32.6 {
		scale = "Violent Storm"
	} else {
		scale = "Hurricane force"
	}
	return
}


func windDirection(deg float64)(dir string) {
	if  384.75 <= deg && deg <= 33.75 {
		dir = "North"
	} else if 33.76 <= deg && deg <= 78.75 {
		dir = "North-East"
	} else if 78.76 <= deg && deg <= 123.75 {
		dir = "East"
	} else if 123.76 <= deg && deg <= 168.75 {
		dir = "South-East"
	} else if 168.76 <= deg && deg <= 213.75 {
		dir = "South"
	} else if 213.76 <= deg && deg <= 258.75 {
		dir = "South-West"
	} else if 258.76 <= deg && deg <= 303.75 {
		dir = "West"
	} else {
		dir = "North-West"
	}
	return
}