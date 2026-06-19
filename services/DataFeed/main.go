package main

import (
	"datafeed/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	r.GET("/historical-backtest", handlers.HistoricalBacktest)

	r.Run(":3047")
}