package test

import (
	"encoding/json"
	"errors"
	"fmt"
	// "net/http"
	// "net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"WheatherAPI/utils"
	"WheatherAPI/tests/mocks"
	_ "WheatherAPI/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase(
		"default",
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8",
			beego.AppConfig.String("mysqluser"),
			beego.AppConfig.String("mysqlpswd"),
			os.Getenv("DB_HOST"),
			beego.AppConfig.String("mysqldb"),
		),
	)
	if err != nil {
		fmt.Printf("%v", err)
	}
}


type MockRequest struct {
	City, Country string
}

func (mock MockRequest) WeatherAPIRequest()(map[string]interface{}, error){
	if mock.City != "Bogota" || mock.Country != "co" {
		return nil, errors.New("City not found")
	}
	response := make(map[string] interface{})
	json.Unmarshal([]byte(mocks.WeatherResponse), &response)
	return response, nil
}


func TestWeatherApi(t *testing.T) {
	m := MockRequest{
		City: "Bogota",
		Country: "co",
	}

	jsonResponse, _ := utils.WeatherReporter(m)
	Convey("Test Weather Endpoint", t, func(){
		Convey("Body must contain all fields", func(){
			So(
				jsonResponse["location_name"] != nil &&
				jsonResponse["temperature"] != nil &&
				jsonResponse["wind"] != nil &&
				jsonResponse["pressure"] != nil &&
				jsonResponse["humidity"] != nil &&
				jsonResponse["geo_coordinates"] != nil &&
				jsonResponse["sunrise"] != nil &&
				jsonResponse["sunset"] != nil &&
				jsonResponse["requested_time"] != nil,
				ShouldBeTrue,
			)
		})
	})
}
