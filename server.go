package main

import (
	"github.com/flaviojmendes/weathergo/config"
	"github.com/flaviojmendes/weathergo/service"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
)

func Server(configuration *config.Configuration, ch *cache.Cache) *gin.Engine {
	router := getRouter()
	router.GET("/health", HealthCheck)
	router.GET("/weather/:lat/:lon/:provider", func(c *gin.Context) {GetWeather(c, ch, configuration)})
	router.Run(configuration.Port)
	log.Print("Server Started on port 8000")
	return router
}

func getRouter() *gin.Engine{
	return gin.Default()
}

func GetWeather(c *gin.Context, ch *cache.Cache, config *config.Configuration) {
	weather, err := service.GetWeather(c.Param("lat"), c.Param("lon"), c.Param("provider"), ch, config)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, weather)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK,"I'm Alive!")
}