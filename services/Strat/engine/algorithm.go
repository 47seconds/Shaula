package engine

import (
	"log"
	"time"

	"strat/models"
)

func ProcessCandle(candle models.Candle) {
	log.Printf(
		"%s %s O:%.2f H:%.2f L:%.2f C:%.2f V:%.2f",
		candle.Symbol,
		candle.Timestamp.Format(time.RFC3339),
		candle.Open,
		candle.High,
		candle.Low,
		candle.Close,
		candle.Volume,
	)

	// SMA
	// EMA
	// RSI
	// Signal generation
	// Paper orders
}