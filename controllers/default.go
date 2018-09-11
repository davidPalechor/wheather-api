package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
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

type jsonResponse struct {
	Name string `json:"name"`
}

func (c *ApiController) Get() {
	var jsonResponse map[string]interface{}
	req := httplib.Get("http://api.openweathermap.org/data/2.5/weather?q=Bogota,co&appid=1508a9a4840a5574c822d70ca2132032")
	data, _ := req.String()

	err := json.Unmarshal([]byte(data), &jsonResponse)

	if err != nil{
		c.Data["json"] = `{"error": 0}`
	}

	c.Data["json"] = jsonResponse
	c.ServeJSON()
}