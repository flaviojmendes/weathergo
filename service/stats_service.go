package service

import (
	"github.com/flaviojmendes/weathergo/config"
	"github.com/recoilme/slowpoke"
	"strconv"
	"time"
)

func increaseRequestsCount(config *config.Configuration) {
	//store
	layout := "2006-01-02"
	statsDate:= time.Now().UTC().Format(layout)
	key := []byte(statsDate)

	totalRequests := GetRequestsCount(config, statsDate)

	totalRequests += 1

	file := config.DbFile


	slowpoke.Set(file, key, []byte(strconv.FormatInt(totalRequests, 10)))
}

func GetRequestsCount(config *config.Configuration, statsDate string) int64{
	// create database
	file := config.DbFile
	// close all opened database
	defer slowpoke.CloseAll()

	// init key
	key := []byte(statsDate)


	// get
	res, err := slowpoke.Get(file, key)

	if err != nil {
		return 0
	}
	totalRequests,_ := strconv.ParseInt(string(res), 10, 64)

	return totalRequests
}