package indicators

import (
	"strat/models"
)

func SMA(candles []models.Candle) float64 {
	if len(candles) == 0 {
		return 0
	}

	var sum float64

	for _, candle := range candles {
		sum += candle.Close
	}

	return sum / float64(len(candles))
}
