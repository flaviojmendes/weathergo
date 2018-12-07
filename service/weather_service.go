package service

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/flaviojmendes/weathergo/config"
	"github.com/flaviojmendes/weathergo/entity"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func GetWeather(c *gin.Context) {
	lat := c.Param("lat")
	lon := c.Param("lon")

	fmt.Println("Retrieving Weather for Latitude: ", lat ," - Longitude", lon)

	key := config.GetConfig().OpenWeatherKeys[rand.Intn(len(config.GetConfig().OpenWeatherKeys))]
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?units=metric&lat=" + lat + "&lon=" + lon + "&appid=" + key)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		v,_ := jason.NewObjectFromReader(response.Body)
		name,_ := v.GetString("name")
		temp,_ := v.GetFloat64("main", "temp")
		hum,_ := v.GetFloat64("main", "humidity")
		weather := entity.Weather{lat,lon,temp,name,hum,time.Now()}
		c.JSON(http.StatusOK, weather)
	}
}