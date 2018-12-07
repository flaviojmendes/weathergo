package main

import (
	"github.com/flaviojmendes/weathergo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Server() {
	router := gin.Default()

	router.GET("/health", HealthCheck)
	router.GET("/weather/:lat/:lon", service.GetWeather)
	router.Run(":8000")

	log.Print("Server Started on port 8000")
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK,"I'm Alive!")
}