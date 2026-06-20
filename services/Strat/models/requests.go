package models

type HistoricalRequest struct {
	Symbol    string `json:"symbol" binding:"required"`
	Timeframe string `json:"timeframe" binding:"required"`
	Rate     int    `json:"rate" binding:"required"`
}
