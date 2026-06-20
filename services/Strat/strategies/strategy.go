package strategies

import "strat/models"

// Strategy interface for every strategy to implement
type Strategy interface {
	Name() string
	OnCandle(
		candle models.Candle,
		candles []models.Candle,
	) *models.Signal
}