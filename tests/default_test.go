package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	"encoding/json"
	_ "WheatherAPI/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}


// TestBeego is a sample to run an endpoint test
func TestWeatherApi(t *testing.T) {
	r, _ := http.NewRequest("GET", "/weather?city=Bogota&country=co", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var jsonResponse map[string] interface{}

	err := json.Unmarshal([]byte(w.Body.String()), &jsonResponse)

	Convey("Test Weather Endpoint", t, func(){
		Convey("Status should be 200", func(){
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Body must be a valid JSON", func(){
			So(err == nil, ShouldBeTrue)
		})
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

