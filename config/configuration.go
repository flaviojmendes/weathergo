package config

import (
	"github.com/tkanos/gonfig"
	"log"
	"os"
)

type Configuration struct {
	Port              	string
	CacheExp			int64
	CachePurge			int64
	OpenWeatherKeys 	[]string
}

func GetConfig() *Configuration{
	configFilePath := os.Getenv("CONFIG_FILE")
	return readConfigFile(configFilePath)
}

func readConfigFile(path string) *Configuration {
	verifyFile(path)
	configuration := Configuration{}
	gonfig.GetConf(path, &configuration)
	return &configuration
}

func verifyFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Panicf("The file %v doesn't exist.", path)
	}
}


