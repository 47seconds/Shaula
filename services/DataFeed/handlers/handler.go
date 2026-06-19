package handlers

import (
	"github.com/gin-gonic/gin"

	"datafeed/models"
	"datafeed/utils"
)

func HealthCheck(c *gin.Context) {
	resp := utils.SuccessResponse(
		"DataFeed service is healthy",
		nil,
	)

	c.JSON(200, resp)
}

func HistoricalBacktest(c *gin.Context) {
	var req models.HistoricalBacktestRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, utils.ErrorResponse(400, err.Error()))
		return
	}

	if req.Symbol == "" {
		req.Symbol = "ADANIENT"
	}

	if req.Timeframe == "" {
		req.Timeframe = "1m"
	}

	if req.Speed == 0 {
		req.Speed = 60
	}

	respData := models.HistoricalBacktestResponse{
		Symbol:    req.Symbol,
		Timeframe: req.Timeframe,
		Speed:     req.Speed,
		Status:    "running",
	}

	resp := utils.SuccessResponse(
		"Historical backtest started",
		respData,
	)

	c.JSON(200, resp)
}
