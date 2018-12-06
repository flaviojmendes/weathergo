package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	Port              int
	OpenWeatherKeys []string
}

func GetConfig() Configuration{
	configuration := Configuration{}
	err := gonfig.GetConf("config.json", &configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}

