package service

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/flaviojmendes/weathergo/config"
	"github.com/flaviojmendes/weathergo/entity"
	"math/rand"
	"net/http"
	"time"
)

func getCurrentWeatherOpenWeather(config *config.Configuration, lat string, lon string) (entity.Weather, error) {
	key := config.OpenWeatherKeys[rand.Intn(len(config.OpenWeatherKeys))]
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?units=metric&lat=" + lat + "&lon=" + lon + "&appid=" + key)
	weather := entity.Weather{}

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		v,_ := jason.NewObjectFromReader(response.Body)

		weather.Humidity,err = v.GetFloat64("main", "humidity")
		weather.Temp,err = v.GetFloat64("main", "temp")
		weather.Location,err = v.GetString("name")
		weather.Sunrise,err = v.GetInt64("sys", "sunrise")
		weather.Sunset,err = v.GetInt64("sys", "sunset")
		weather.Lat = lat
		weather.Lon = lon
		weather.RetrievedAt = time.Now()

	}
	return weather,err

}


