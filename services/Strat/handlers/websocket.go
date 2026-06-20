package handlers

import (
	"net/http"
	"log"

	"strat/utils"
	"strat/models"
	"strat/engine"

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

	rate := utils.StringToInt(
		c.DefaultQuery("rate", "60"),
		60,
	)

	log.Printf("Received WebSocket request for symbol: %s, timeframe: %s, rate: %d seconds\n", symbol, timeframe, rate)

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
		"rate":     rate,
	})

	eng := engine.New()

	for {
		var msg models.CandleMessage

		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsCloseError(
				err,
				websocket.CloseNormalClosure,
				websocket.CloseGoingAway,
			) {
				log.Printf(
					"Backtest completed: %s %s",
					symbol,
					timeframe,
				)
			} else {
				log.Printf("WS error: %v", err)
			}

			break
		}

		if msg.Type != "candle" {
			continue
		}

		eng.ProcessCandle(msg.Data)
	}

	log.Printf(
			"Buffer size: %d",
			eng.Count(),
		)
}
