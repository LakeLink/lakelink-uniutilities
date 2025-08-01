package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetWeatherHandler handles the weather API request
func GetWeatherHandler(c *gin.Context) {
	geopos := c.Query("geopos")
	if geopos == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "geopos parameter is required"})
		return
	}

	caiyunWeatherToken := os.Getenv("CAIYUN_WEATHER_TOKEN")
	if caiyunWeatherToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CAIYUN_WEATHER_TOKEN not set"})
		return
	}

	caiyunURL := fmt.Sprintf("https://api.caiyunapp.com/v2.6/%s/%s/weather?alert=true&dailysteps=1&hourlysteps=24", caiyunWeatherToken, geopos)

	log.Printf("Requesting weather data from Caiyun API: %s", caiyunURL)

	resp, err := http.Get(caiyunURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch weather data: %v", err)})
		log.Printf("Failed to fetch weather data: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read response body: %v", err)})
		log.Printf("Failed to read response body: %v", err)
		return
	}

	var caiyunResp CaiyunAPIResponse
	err = json.Unmarshal(body, &caiyunResp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to parse weather data: %v", err)})
		log.Printf("Failed to parse weather data: %v", err)
		return
	}

	if caiyunResp.Status != "ok" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Caiyun API returned status: %s; msg=%s", caiyunResp.Status, caiyunResp.ErrorMsg)})
		log.Printf("Caiyun API returned status: %s", caiyunResp.Status)
		return
	}

	weatherData := gin.H{
		"realtime": caiyunResp.Result.Realtime,
		"alert":    caiyunResp.Result.Alert,
		// "minutely": caiyunResp.Result.Minutely,
		"hourly": caiyunResp.Result.Hourly,
		"daily":  caiyunResp.Result.Daily,
	}

	c.JSON(http.StatusOK, weatherData)
}

// HelloHandler handles the root endpoint
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, world",
	})
}
