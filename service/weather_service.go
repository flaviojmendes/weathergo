package service

import (
	"encoding/json"
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/flaviojmendes/weathergo/config"
	"github.com/flaviojmendes/weathergo/entity"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lat := params["lat"]
	lon := params["lon"]

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
		json.NewEncoder(w).Encode(weather)



	}
}