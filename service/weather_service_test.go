package service

import (
	"github.com/flaviojmendes/weathergo/config"
	"github.com/patrickmn/go-cache"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGetWeather(t *testing.T) {
	ch := cache.New(time.Duration(1)*time.Minute, time.Duration(1)*time.Minute)
	configFile := &config.Configuration{
		OpenWeatherKeys: []string{"1455382c9be6c3db4fe8f894230202b7"},
	}

	convey.Convey("get weather to 52.0984794,-9.7957126 in OPENWEATHER should return Killorglin", t, func() {
		weather,_ := GetWeather("52.0984794", "-9.7957126","OPENWEATHER",ch,configFile)
		convey.So(weather.Location, convey.ShouldEqual, "Killorglin")

		weatherCache,_ := GetWeather("52.0984794", "-9.7957126","OPENWEATHER",ch,configFile)
		convey.So(weatherCache.Location, convey.ShouldEqual, weather.Location)
	})
}

func TestGetWeatherWithInvalidKey(t *testing.T) {
	ch := cache.New(time.Duration(1)*time.Minute, time.Duration(1)*time.Minute)
	configFile := &config.Configuration{
		OpenWeatherKeys: []string{"1455382c9be6c3db4fe8f894230202b7_invalid"},
	}

	convey.Convey("get weather to 52.0984794,-9.7957126 in OPENWEATHER should return Killorglin", t, func() {
		_,err := GetWeather("52.0984794", "-9.7957126","OPENWEATHER",ch,configFile)
		convey.So(err, convey.ShouldNotBeNil)
	})
}