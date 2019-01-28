package service

import (
	"errors"
	"fmt"
	"github.com/flaviojmendes/weathergo/config"
	"github.com/flaviojmendes/weathergo/entity"
	"github.com/patrickmn/go-cache"
	"strings"
)

func GetWeather(lat string, lon string, provider string, ch *cache.Cache, config *config.Configuration) (entity.Weather, error) {

	var response entity.Weather
	var err error

	fmt.Println("Retrieving Weather for Latitude: ", lat ," - Longitude", lon, " - Provider: ", provider)

	cacheKey := strings.Join([]string{lat,lon,provider},"-")
	cachedResponse, found := ch.Get(cacheKey)

	if found {
		fmt.Println("Found in cache. Returning...")
		response := cachedResponse.(entity.Weather)
		return response,err
	}

	switch provider {
	case "OPENWEATHER":
		response, err = getCurrentWeatherOpenWeather(config, lat,lon)
		ch.Set(cacheKey, response, cache.DefaultExpiration)
	default:
		err = errors.New("unfortunately we are just supporting OPENWEATHER provider")
	}

	increaseRequestsCount(config)

	return response,err
}