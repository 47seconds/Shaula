package handlers

import (
	// "strat/models"
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
