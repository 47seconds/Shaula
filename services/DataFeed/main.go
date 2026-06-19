package main

import (
	"datafeed/handlers"
	"datafeed/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	r.POST("/historical", handlers.Historical)

	r.Run(":" + utils.GetEnvString("DATAFEED_PORT", "3047"))
}