package main

import (
	"github.com/flaviojmendes/weathergo/config"
	"github.com/flaviojmendes/weathergo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Server(configuration *config.Configuration) {
	router := gin.Default()

	router.GET("/health", HealthCheck)
	router.GET("/weather/:lat/:lon", service.GetWeather)
	router.Run(configuration.Port)

	log.Print("Server Started on port 8000")
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK,"I'm Alive!")
}