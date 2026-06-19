package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func HistoricalBacktest(c *gin.Context) {
	symbol := c.DefaultQuery("symbol", "ADANIENT")
	timeframe := c.DefaultQuery("timeframe", "1m")

	speedStr := c.DefaultQuery("speed", "60")
	speed, err := strconv.Atoi(speedStr)
	if err != nil {
		speed = 60
	}

	c.JSON(200, gin.H{
		"symbol":    symbol,
		"timeframe": timeframe,
		"speed":     speed,
	})
}