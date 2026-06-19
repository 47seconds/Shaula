package main

import (
	"log"
	"strat/handlers"
	"strat/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	r.GET("/historical-ws", handlers.HistoricalWS)

	if err := r.Run(":" + utils.GetEnvString("STRAT_PORT", "3048")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
