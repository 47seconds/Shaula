package strategies

import (
	"strat/indicators"
	"strat/models"
)

type SMA20 struct{}

func (s *SMA20) Name() string {
	return "sma20"
}

func (s *SMA20) OnCandle(
	candle models.Candle,
	candles []models.Candle,
) *models.Signal {

	last20 := candles

	if len(last20) > 20 {
		last20 = last20[len(last20)-20:]
	}

	if len(last20) < 20 {
		return nil
	}

	sma20 := indicators.SMA(last20)

	signal := &models.Signal{
		Symbol:    candle.Symbol,
		Timeframe: candle.Timeframe,
		Price:     candle.Close,
		Timestamp: candle.Timestamp,
		Type:      models.Hold,
	}

	if candle.Close > sma20 {
		signal.Type = models.Buy
	} else if candle.Close < sma20 {
		signal.Type = models.Sell
	}

	return signal
}
