package main

import (
	"log"

	"datafeed/db"
	"datafeed/handlers"
	"datafeed/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("Connected to PostgreSQL")

	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	r.POST("/historical", handlers.Historical)

	if err := r.Run(":" + utils.GetEnvString("DATAFEED_PORT", "3047")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}