package main

import config "github.com/flaviojmendes/weathergo/config"

func main() {
	configuration := config.GetConfig()

	Server(configuration)
}
