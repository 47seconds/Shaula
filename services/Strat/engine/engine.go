package engine

import (
	"log"
	// "time"

	"strat/models"
	"strat/utils"
	// "strat/indicators"
	"strat/strategies"
)

type Engine struct {
	Candles []models.Candle
	MaxSize int
	Strategies []strategies.Strategy
}

func New(strats ...strategies.Strategy) *Engine {
	return &Engine{
		MaxSize:    utils.MAX_CANDLES_SIZE,
		Strategies: strats,
	}
}

func (e *Engine) Count() int {
	return len(e.Candles)
}

func (e *Engine) ProcessCandle(candle models.Candle) {
	e.Candles = append(e.Candles, candle)

	if len(e.Candles) > e.MaxSize {
		e.Candles = e.Candles[1:]
	}

	for _, strategy := range e.Strategies {
		signal := strategy.OnCandle(
			candle,
			e.Candles,
		)

		if signal == nil {
			continue
		}

		log.Printf(
			"[%s] %s %.2f",
			strategy.Name(),
			signal.Type,
			signal.Price,
		)
	}
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
