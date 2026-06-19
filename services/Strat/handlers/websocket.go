package handlers

import (
	"strat/utils"

	"github.com/gin-gonic/gin"
	// "github.com/gorilla/websocket"
)

func HistoricalWS(c *gin.Context) {
	symbol := c.Query("symbol")

	if symbol == "" {
		c.JSON(
			utils.BAD_REQUEST,
			utils.ErrorResponse(
				utils.BAD_REQUEST,
				"symbol is required",
			),
		)
		return
	}

	timeframe := c.DefaultQuery("timeframe", "1m")
	
	speed := utils.StringToInt(
		c.DefaultQuery("speed", "60"),
		60,
	)

	resp := utils.SuccessResponse(
		"WebSocket request validated",
		gin.H{
			"symbol":    symbol,
			"timeframe": timeframe,
			"speed":     speed,
		},
	)

	c.JSON(utils.OK, resp)
}