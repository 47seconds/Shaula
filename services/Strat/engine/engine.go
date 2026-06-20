package engine

import (
	"log"
	// "time"

	"strat/models"
	"strat/utils"
	"strat/indicators"
)

type Engine struct {
	Candles []models.Candle
	MaxSize int
}

func New() *Engine {
	return &Engine{
		MaxSize: utils.MAX_CANDLES_SIZE,
	}
}

func (e *Engine) Count() int {
	return len(e.Candles)
}

func (e *Engine) ProcessCandle(candle models.Candle) {
	// log.Printf(
	// 	"%s %s O:%.2f H:%.2f L:%.2f C:%.2f V:%.2f",
	// 	candle.Symbol,
	// 	candle.Timestamp.Format(time.RFC3339),
	// 	candle.Open,
	// 	candle.High,
	// 	candle.Low,
	// 	candle.Close,
	// 	candle.Volume,
	// )

	e.Candles = append(e.Candles, candle)

	if len(e.Candles) > e.MaxSize {
		e.Candles = e.Candles[1:]
	}

	// Prototype: Calculate SMA20
	last20 := e.Last(20)

	if len(last20) == 20 {
		sma20 := indicators.SMA(last20)

		log.Printf(
			"Close: %.2f SMA20: %.2f",
			candle.Close,
			sma20,
		)
	}

	// SMA
	// EMA
	// RSI
	// Signal generation
	// Paper orders
}

func (e *Engine) Last(n int) []models.Candle {
	if n <= 0 {
		return nil
	}

	if n > len(e.Candles) {
		n = len(e.Candles)
	}

	return e.Candles[len(e.Candles)-n:]
}
