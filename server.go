package main

import (
	"github.com/flaviojmendes/weathergo/config"
	"github.com/flaviojmendes/weathergo/service"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/pjebs/restgate"
	"log"
	"net/http"
)

func Server(configuration *config.Configuration, ch *cache.Cache) *gin.Engine {
	router := getRouter()

	// Initialize Restgate
	rg := restgate.New("X-Auth-Key", "X-Auth-Secret", restgate.Static, restgate.Config{
		Key:                []string{configuration.AuthKey},
		Secret:             []string{configuration.AuthSecret},
		HTTPSProtectionOff: true,
	})
	// Create Gin middleware - integrate Restgate with Gin
	rgAdapter := func(c *gin.Context) {
		nextCalled := false
		nextAdapter := func(http.ResponseWriter, *http.Request) {
			nextCalled = true
			c.Next()
		}
		rg.ServeHTTP(c.Writer, c.Request, nextAdapter)
		if nextCalled == false {
			c.AbortWithStatus(401)
		}
	}

	// Use Restgate with Gin
	router.Use(rgAdapter)

	router.GET("/health", HealthCheck)
	router.GET("/wea	ther/:lat/:lon/:provider", func(c *gin.Context) {GetWeather(c, ch, configuration)})
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