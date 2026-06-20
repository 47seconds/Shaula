package handlers

import (
	"log"
	"time"
	"io"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v4"

	"datafeed/db"
	"datafeed/models"
	"datafeed/utils"
	"datafeed/feeds"
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

	if req.Rate < 0 {
		c.JSON(
			utils.BAD_REQUEST,
			utils.ErrorResponse(
				utils.BAD_REQUEST,
				"rate must be non-negative",
			),
		)
		return
	} else if req.Rate == 0 {
		req.Rate = 1
	}

	if req.Timeframe == "" {
		req.Timeframe = "1m"
	}

	if req.Candles < 0 {
		c.JSON(
			utils.BAD_REQUEST,
			utils.ErrorResponse(
				utils.BAD_REQUEST,
				"number of candles must be non-negative, 0 for all",
			),
		)
		return
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

		var rows pgx.Rows

		if req.Candles > 0 {
			rows, err = db.Pool.Query(
				context.Background(),
				utils.DB_QUERY_LIMIT,
				req.Symbol,
				req.Timeframe,
				req.Candles,
			)
		} else {
			rows, err = db.Pool.Query(
				context.Background(),
				utils.DB_QUERY_ALL,
				req.Symbol,
				req.Timeframe,
			)
		}
		if err != nil {
			log.Printf("Query failed: %v", err)
			return
		}

		feed := feeds.NewPGFeed(rows)
		defer feed.Close()

		for {
			candle, err := feed.Next()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Feed error: %v", err)
				return
			}

			msg := models.CandleMessage{
				Type: "candle",
				Data: *candle,
			}

			if err := conn.WriteJSON(msg); err != nil {
				log.Printf("Failed to send candle: %v", err)
				return
			}

			time.Sleep(time.Second / time.Duration(req.Rate))
		}
	}()

	respData := models.HistoricalResponse{
		Symbol:    req.Symbol,
		Timeframe: req.Timeframe,
		Rate:     req.Rate,
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