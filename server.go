package main

import (
	"github.com/flaviojmendes/weathergo/config"
	"github.com/flaviojmendes/weathergo/service"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
)

func Server(configuration *config.Configuration, ch *cache.Cache) {
	router := gin.Default()

	router.GET("/health", HealthCheck)
	router.GET("/weather/:lat/:lon/:provider", func(c *gin.Context) {service.GetWeather(c, ch)})
	router.Run(configuration.Port)

	log.Print("Server Started on port 8000")
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK,"I'm Alive!")
}