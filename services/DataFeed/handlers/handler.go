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

	c.JSON(utils.OK, resp)
}

func HistoricalBacktest(c *gin.Context) {
	var req models.HistoricalBacktestRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(utils.BAD_REQUEST, utils.ErrorResponse(utils.BAD_REQUEST, err.Error()))
		return
	}

	errMsg := ""

	if req.Speed < 0 {
		errMsg += "Speed must be non-negative."
	}

	if req.Symbol == "" {
		errMsg += "Symbol is required."
	}

	// These are default values, so we don't need to validate them
	if req.Timeframe == "" {
		req.Timeframe = "1m"
	}

	if req.Speed == 0 {
		req.Speed = 60
	}

	if errMsg != "" {
		c.JSON(utils.BAD_REQUEST, utils.ErrorResponse(utils.BAD_REQUEST, errMsg))
		return
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

	c.JSON(utils.OK, resp)
}
