package mocks


const WeatherResponse string = `{
	"coord": {
		"lon": -74.08,
		"lat": 4.6
	},
	"weather": [
		{
			"id": 803,
			"main": "Clouds",
			"description": "broken clouds",
			"icon": "04d"
		}
	],
	"base": "stations",
	"main": {
		"temp": 292.15,
		"pressure": 1025,
		"humidity": 42,
		"temp_min": 292.15,
		"temp_max": 292.15
	},
	"visibility": 10000,
	"wind": {
		"speed": 6.2,
		"deg": 130
	},
	"clouds": {
		"all":75
	},
	"dt": 1537552800,
	"sys": {
		"type": 1,
		"id": 4245,
		"message": 0.0085,
		"country":"CO",
		"sunrise":1537526743,
		"sunset":1537570342
	},
	"id": 3688689,
	"name": "Bogota",
	"cod":200
}
`