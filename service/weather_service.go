package service

import (
	"errors"
	"fmt"
	"github.com/flaviojmendes/weathergo/entity"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
	"strings"
)

func GetWeather(c *gin.Context, ch *cache.Cache) {
	lat := c.Param("lat")
	lon := c.Param("lon")
	provider := c.Param("provider")

	var response entity.Weather
	var err error

	fmt.Println("Retrieving Weather for Latitude: ", lat ," - Longitude", lon, " - Provider: ", provider)

	cacheKey := strings.Join([]string{lat,lon,provider},"-")
	cachedResponse, found := ch.Get(cacheKey)

	if found {
		fmt.Println("Found in cache. Returning...")
		c.JSON(http.StatusOK, cachedResponse)
		return
	}

	switch provider {
	case "OPENWEATHER":
		response, err = getCurrentWeatherOpenWeather(lat,lon)
		ch.Set(cacheKey, &response, cache.DefaultExpiration)
	default:
		err = errors.New("unfortunately we are just supporting OPENWEATHER provider")
	}


	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,err)
	} else {
		c.JSON(http.StatusOK, response)
	}
}