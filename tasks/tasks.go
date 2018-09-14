package tasks

import (
	"github.com/astaxie/beego/toolbox"
	"WheatherAPI/utils"
)

func WeatherTask(city string, country string) error {
	task := toolbox.NewTask("weatherTask", "* * 0/1 * * *", func() error{
		utils.WeatherReporter(city, country)
		return nil
	})
	toolbox.AddTask("weatherTask", task)
	toolbox.StartTask()
	return nil
}