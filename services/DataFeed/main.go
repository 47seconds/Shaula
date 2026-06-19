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

	if err := r.Run(":" + utils.GetEnvString("DATAFEED_PORT", "3047")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}