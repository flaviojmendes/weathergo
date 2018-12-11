package main

import (
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersResource(t *testing.T) {
	router := getRouter()
	router.GET("/health", HealthCheck)
	convey.Convey("GET request to /health should return 200", t, func() {
		req, _ := http.NewRequest("GET", "/health", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		convey.So(resp.Code, convey.ShouldEqual, http.StatusOK)
	})
}