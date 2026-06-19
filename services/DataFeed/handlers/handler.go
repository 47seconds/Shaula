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

func Historical(c *gin.Context) {
	var req models.HistoricalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := ""

		if err.Error() == "EOF" {
			errMsg = "request body is required"
		} else {
			errMsg = err.Error()
		}

		c.JSON(
			utils.BAD_REQUEST,
			utils.ErrorResponse(
				utils.BAD_REQUEST,
				errMsg,
			),
		)
		return
	}

	errMsg := ""

	if req.Speed < 0 {
		errMsg += "Speed must be non-negative"
	}

	if req.Timeframe == "" {
		req.Timeframe = "1m"
	}

	if req.Speed == 0 {
		req.Speed = 60
	}

	if errMsg != "" {
		c.JSON(
			utils.BAD_REQUEST,
			utils.ErrorResponse(
				utils.BAD_REQUEST,
				errMsg,
			),
		)
		return
	}

	respData := models.HistoricalResponse{
		Symbol:    req.Symbol,
		Timeframe: req.Timeframe,
		Speed:     req.Speed,
		Status:    "running",
	}

	c.JSON(
		utils.OK,
		utils.SuccessResponse(
			"Historical backtest started",
			respData,
		),
	)
}
