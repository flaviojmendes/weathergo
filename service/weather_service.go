package service

import (
	"errors"
	"fmt"
	"github.com/flaviojmendes/weathergo/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWeather(c *gin.Context) {
	lat := c.Param("lat")
	lon := c.Param("lon")
	provider := c.Param("provider")

	fmt.Println("Retrieving Weather for Latitude: ", lat ," - Longitude", lon, " - Provider: ", provider)

	var response entity.Weather
	var err error

	switch provider {
	case "OPENWEATHER":
		response, err = getCurrentWeatherOpenWeather(lat,lon)
	default:
		err = errors.New("unfortunately we are just supporting OPENWEATHER provider")
	}


	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,err)
	} else {
		c.JSON(http.StatusOK, response)
	}
}