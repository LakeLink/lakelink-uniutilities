package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	LoadConfig()

	r := gin.Default()

	r.SetTrustedProxies(nil)

	setupRoutes(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
