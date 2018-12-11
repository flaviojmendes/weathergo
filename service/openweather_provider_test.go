package service

import (
	"github.com/flaviojmendes/weathergo/config"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetCurrentWeatherOpenWeather(t *testing.T) {

	configFile := &config.Configuration{
		OpenWeatherKeys: []string{"1455382c9be6c3db4fe8f894230202b7"},
	}

	convey.Convey("get openweather weather", t, func() {
		w,_ := getCurrentWeatherOpenWeather(configFile, "52.0984794","-9.7957126")
		convey.So("Killorglin", convey.ShouldEqual, w.Location)
	})
}
