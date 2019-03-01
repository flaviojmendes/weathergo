package main

import (
	"github.com/flaviojmendes/weathergo/config"
	"github.com/flaviojmendes/weathergo/service"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/semihalev/gin-stats"
	"log"
	. "net/http"
)

func server(configuration *config.Configuration, ch *cache.Cache) *gin.Engine {
	router := getRouter()


	router.Use(stats.RequestStats())

	router.GET("/stats", func(c *gin.Context) {
		c.JSON(StatusOK, stats.Report())
	})


	router.GET("/health", HealthCheck)
	router.GET("/weather/:lat/:lon/:provider", func(c *gin.Context) {GetWeather(c, ch, configuration)})

	if gin.IsDebugging() {
		router.Run(configuration.DebugPort)
	} else {
		log.Fatal(autotls.Run(router, configuration.WhiteListHosts...))
	}
	return router
}

func getRouter() *gin.Engine{
	return gin.Default()
}

func GetWeather(c *gin.Context, ch *cache.Cache, config *config.Configuration) {
	weather, err := service.GetWeather(c.Param("lat"), c.Param("lon"), c.Param("provider"), ch, config)
	if err != nil {
		c.AbortWithError(StatusInternalServerError, err)
	}
	c.JSON(StatusOK, weather)
}

func HealthCheck(c *gin.Context) {
	c.JSON(StatusOK,"I'm Alive!")
}