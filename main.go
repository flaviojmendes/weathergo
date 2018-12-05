package main

import (
	"encoding/json"
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/flaviojmendes/weathergo/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const apiKey = "1455382c9be6c3db4fe8f894230202b7"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/weather/{lat}/{lon}", GetWeather).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lat := params["lat"]
	lon := params["lon"]
	fmt.Println("Retrieving Weather for Latitude: ", lat ," - Longitude", lon)
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?units=metric&lat=" + lat + "&lon=" + lon + "&appid=" + apiKey)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		v,_ := jason.NewObjectFromReader(response.Body)
		name,_ := v.GetString("name")
		temp,_ := v.GetFloat64("main", "temp")
		hum,_ := v.GetString("main", "humidity")
		weather := entity.Weather{lat,lon,temp,name,hum,time.Now()}
		json.NewEncoder(w).Encode(weather)



	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("I'm Alive!")
}
