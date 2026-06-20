package handlers

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

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

	if req.Speed < 0 {
		c.JSON(
			utils.BAD_REQUEST,
			utils.ErrorResponse(
				utils.BAD_REQUEST,
				"speed must be non-negative",
			),
		)
		return
	}

	if req.Timeframe == "" {
		req.Timeframe = "1m"
	}

	if req.Speed == 0 {
		req.Speed = 60
	}

	// Start streaming in background
	go func() {
		stratHost := utils.GetEnvString("STRAT_HOST", "localhost")
		stratPort := utils.GetEnvString("STRAT_PORT", "3048")

		wsURL := "ws://" + stratHost + ":" + stratPort + "/historical-ws" +
			"?symbol=" + req.Symbol +
			"&timeframe=" + req.Timeframe

		conn, _, err := websocket.DefaultDialer.Dial(
			wsURL,
			nil,
		)
		if err != nil {
			log.Printf("Failed to connect to Strat WS: %v", err)
			return
		}
		defer func() {
			conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(
					websocket.CloseNormalClosure,
					"backtest completed",
				),
			)

			conn.Close()
		}()

		for i := 0; i < 10; i++ {
			candle := models.Candle{
				Symbol:    req.Symbol,
				Timeframe: req.Timeframe,
				Timestamp: time.Now(),

				Open:   100 + float64(i),
				High:   105 + float64(i),
				Low:    95 + float64(i),
				Close:  102 + float64(i),

				Volume: 1000,
			}

			msg := models.CandleMessage{
				Type: "candle",
				Data: candle,
			}

			if err := conn.WriteJSON(msg); err != nil {
				log.Printf("Failed to send candle: %v", err)
				return
			}

			log.Printf(
				"Sent candle %d: %s %.2f",
				i,
				candle.Symbol,
				candle.Close,
			)

			time.Sleep(time.Second)
		}
	}()

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