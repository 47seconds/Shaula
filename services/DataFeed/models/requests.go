package models

type HistoricalBacktestRequest struct {
	Symbol    string `json:"symbol"`
	Timeframe string `json:"timeframe,omitempty"`
	Speed     int    `json:"speed,omitempty"`
}
