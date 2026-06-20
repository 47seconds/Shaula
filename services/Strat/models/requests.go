package models

type HistoricalRequest struct {
	Symbol    string `json:"symbol" binding:"required"`
	Timeframe string `json:"timeframe" binding:"required"`
	Speed     int    `json:"speed" binding:"required"`
}
