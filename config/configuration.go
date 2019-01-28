package config

import (
	"github.com/tkanos/gonfig"
	"log"
	"os"
)

type Configuration struct {
	DebugPort           string
	CacheExp			int64
	CachePurge			int64
	OpenWeatherKeys 	[]string
	WhiteListHosts		[]string
	DbFile				string
}

func GetConfig() *Configuration{
	configFilePath := os.Getenv("CONFIG_FILE")
	log.Printf("Getting config file from %s",configFilePath)
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


