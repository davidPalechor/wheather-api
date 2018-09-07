package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
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
	req := httplib.Get("http://api.openweathermap.org/data/2.5/weather?q=Bogota,co&appid=1508a9a4840a5574c822d70ca2132032")
	data, _ := req.String()
	fmt.Printf("%q\n", data)
	jsonResponse := jsonResponse{"Sergio David"}
	c.Data["json"] = &jsonResponse
	c.ServeJSON()
}