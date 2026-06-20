package models

type HistoricalRequest struct {
	Symbol    string `json:"symbol" binding:"required"`
	Timeframe string `json:"timeframe"`
	Rate     int    `json:"rate"`
	Candles   int    `json:"candles"` // 0 = all
}
