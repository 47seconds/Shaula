package handlers

import (
	"strat/models"
	"strat/utils"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	resp := utils.SuccessResponse(
		"Strat service is healthy",
		nil,
	)

	c.JSON(utils.OK, resp)
}

// will use this later when redis and all for generating session_id and returning it to the client for websocket connection
func Historical(c *gin.Context) {
	var req models.HistoricalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			utils.BAD_REQUEST,
			utils.ErrorResponse(
				utils.BAD_REQUEST,
				"Invalid request: " + err.Error(),
			),
		)
		return
	}

	// genrate session_id instead, store the data into memory or redis, and return the session_id to the client for websocket connection
	resp := utils.SuccessResponse(
		"Historical request validated",
		gin.H{
			"symbol":    req.Symbol,
			"timeframe": req.Timeframe,
			"rate":     req.Rate,
		},
	)

	c.JSON(utils.OK, resp)
}
