package handlers

import (
	"net/http"
	"time"
	"log"

	"strat/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

	log.Printf("Received WebSocket request for symbol: %s, timeframe: %s, speed: %d seconds\n", symbol, timeframe, speed)

	conn, err := upgrader.Upgrade(
		c.Writer,
		c.Request,
		nil,
	)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v\n", err)
		return
	}
	defer conn.Close()

	conn.WriteJSON(gin.H{
		"type":      "connected",
		"symbol":    symbol,
		"timeframe": timeframe,
		"speed":     speed,
	})

	log.Printf("WebSocket connection established for symbol: %s, timeframe: %s, speed: %d seconds\n", symbol, timeframe, speed)

	for i := 0; i < 10; i++ {
		conn.WriteJSON(gin.H{
			"type": "candle",
			"data": gin.H{
				"index": i,
				"open":  100 + i,
				"high":  105 + i,
				"low":   95 + i,
				"close": 102 + i,
			},
		})

		time.Sleep(time.Second)
	}
}