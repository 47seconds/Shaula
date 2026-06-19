package main

import (
	"datafeed/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	r.POST("/historical-backtest", handlers.HistoricalBacktest)

	r.Run(":3047")
}