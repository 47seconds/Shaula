package handlers

import (
	"log"
	"net/http"
	"time"

	"strat/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for development
	},
}

func HistoricalWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		return
	}
	defer conn.Close()

	for {
		var candle models.Candle

		err := conn.ReadJSON(&candle)
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		log.Printf(
			"%s %s %.2f",
			candle.Symbol,
			candle.Timestamp.Format(time.RFC3339),
			candle.Close,
		)

		// strategy.Process(candle)
	}
}

