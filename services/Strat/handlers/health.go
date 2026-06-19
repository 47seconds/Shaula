package handlers

import (
	"net/http"
	"strat/models"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.HealthResponse{
		Status:  "ok",
		Message: "Strat service is running",
	})
}
