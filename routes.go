package main

import (
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/", HelloHandler)
	r.GET("/api/weather", GetWeatherHandler)
}
