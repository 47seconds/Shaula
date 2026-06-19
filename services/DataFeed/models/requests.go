package models

type HistoricalRequest struct {
	Symbol    string `json:"symbol" binding:"required"`
	Timeframe string `json:"timeframe"`
	Speed     int    `json:"speed"`
}
