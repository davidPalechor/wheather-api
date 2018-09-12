package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	// "fmt"
)

type Request struct {
	Id 				int
	LocationName 	string
	Temperature	 	string
	Wind 			string
	Pressure 		string
	Humidity 		string
	Sunrise 		time.Time
	Sunset 			time.Time
	Long			float64
	Lat				float64
	RequestedTime 	time.Time
	Timestamp 		time.Time `orm:"auto_now;type(datetime)"`
}

func (this *Request) TableName() string {
    return "request"
}

func init() {
    // Need to register model in init
	orm.RegisterModel(new(Request))
	// err := orm.RunSyncdb("default", true, true)
	// if err != nil{
	// 	fmt.Println(err)
	// }
}

