package service

import (
	"github.com/flaviojmendes/weathergo/config"
	"github.com/recoilme/slowpoke"
	"strconv"
)

func increaseRequestsCount(config *config.Configuration) {

	totalRequests := GetRequestsCount(config)

	totalRequests += 1

	//store
	file := config.DbFile
	key := []byte("total_requests")

	slowpoke.Set(file, key, []byte(strconv.Itoa(totalRequests)))
}

func GetRequestsCount(config *config.Configuration) int{
	// create database
	file := config.DbFile
	// close all opened database
	defer slowpoke.CloseAll()

	// init key
	key := []byte("total_requests")

	// get
	res, err := slowpoke.Get(file, key)

	if err != nil {
		return 0
	}
	totalRequests,_ := strconv.Atoi(string(res))

	return totalRequests
}