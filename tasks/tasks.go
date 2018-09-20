package tasks

import (
	"github.com/astaxie/beego/toolbox"
	"WheatherAPI/utils"
)

func WeatherTask(city string, country string) error {
	request := utils.Request {
		City:		city,
		Country:	country,
	}
	task := toolbox.NewTask("weatherTask", "* * */1 * * *", func() error{
		request.WeatherReporter()
		return nil
	})
	toolbox.AddTask("weatherTask", task)
	toolbox.StartTask()
	return nil
}