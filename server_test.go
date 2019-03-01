package main

import (
	"github.com/antonholmquist/jason"
	"github.com/flaviojmendes/weathergo/config"
	"github.com/patrickmn/go-cache"
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthCheck(t *testing.T) {
	ch := cache.New(time.Duration(1)*time.Minute, time.Duration(1)*time.Minute)
	configFile := &config.Configuration{
		OpenWeatherKeys: []string{"1455382c9be6c3db4fe8f894230202b7"},
	}
	router := server(configFile, ch)
	convey.Convey("GET request to /health should return 200", t, func() {
		req, _ := http.NewRequest("GET", "/health", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		convey.So(resp.Code, convey.ShouldEqual, http.StatusOK)
	})
}


func TestGetStats(t *testing.T) {
	ch := cache.New(time.Duration(1)*time.Minute, time.Duration(1)*time.Minute)
	configFile := &config.Configuration{
		OpenWeatherKeys: []string{"1455382c9be6c3db4fe8f894230202b7"},
	}
	router := server(configFile, ch)
	convey.Convey("GET request to /stats should return 200", t, func() {
		req, _ := http.NewRequest("GET", "/stats", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		convey.So(resp.Code, convey.ShouldEqual, http.StatusOK)
	})
}

func TestWeatherResource(t *testing.T) {
	ch := cache.New(time.Duration(1)*time.Minute, time.Duration(1)*time.Minute)
	configFile := &config.Configuration{
		OpenWeatherKeys: []string{"1455382c9be6c3db4fe8f894230202b7"},
	}

	router := server(configFile, ch)

	convey.Convey("GET request to /weather/52.0984794/-9.7957126/OPENWEATHER should return Killorglin", t, func() {
		req, _ := http.NewRequest("GET", "/weather/52.0984794/-9.7957126/OPENWEATHER", nil)
		req.Header.Add("X-Auth-Key", "test")
		req.Header.Add("X-Auth-Secret", "test")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		convey.So(resp.Code, convey.ShouldEqual, http.StatusOK)
		v,_ := jason.NewObjectFromReader(resp.Body)
		name,_ := v.GetString("location")
		convey.So(name, convey.ShouldEqual, "Killorglin")
	})
}

func TestWeatherResourceWithInvalidKey(t *testing.T) {
	ch := cache.New(time.Duration(1)*time.Minute, time.Duration(1)*time.Minute)
	configFile := &config.Configuration{
		OpenWeatherKeys: []string{"1455382c9be6c3db4fe8f894230202b7_invalid"},
	}

	router := server(configFile, ch)

	convey.Convey("GET request to /weather/52.0984794/-9.7957126/OPENWEATHER should return Killorglin", t, func() {
		req, _ := http.NewRequest("GET", "/weather/52.0984794/-9.7957126/OPENWEATHER", nil)
		req.Header.Add("X-Auth-Key", "test")
		req.Header.Add("X-Auth-Secret", "test")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		convey.So(resp.Code, convey.ShouldEqual, http.StatusInternalServerError)
	})
}

