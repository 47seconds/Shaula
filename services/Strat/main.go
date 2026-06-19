package main

import (
	"log"
	"strat/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	r.GET("/historical-ws", handlers.HistoricalWS)

	log.Println("Starting Strat service on :3048")
	if err := r.Run(":3048"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
