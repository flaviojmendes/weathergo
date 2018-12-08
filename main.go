package main

import (
	"github.com/flaviojmendes/weathergo/config"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

func main() {
	configuration := config.GetConfig()
	ch := cache.New(time.Duration(configuration.CacheExp)*time.Minute, time.Duration(configuration.CachePurge)*time.Minute)
	log.Printf("Cache set to have %d minutes of duration. It will purge expired data every %d minutes.", configuration.CacheExp, configuration.CachePurge)

	Server(configuration, ch)
}
