package config

import (
	"github.com/tkanos/gonfig"
	"log"
	"os"
)

type Configuration struct {
	Port              string
	OpenWeatherKeys []string
}

func GetConfig() *Configuration{
	configFile := os.Getenv("CONFIG_FILE")
	verifyFile(configFile)

	configuration := Configuration{}
	gonfig.GetConf(configFile, &configuration)

	return &configuration
}


func verifyFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Panicf("The file %v doesn't exist.", path)
	}
}


