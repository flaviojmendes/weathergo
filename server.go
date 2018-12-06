package main

import (
	"encoding/json"
	"github.com/flaviojmendes/weathergo/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Server() {
	router := mux.NewRouter()
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/weather/{lat}/{lon}", service.GetWeather).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("I'm Alive!")
}