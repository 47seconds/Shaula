package models

type Response struct {
	Code	int    `json:"code"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data	any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type HistoricalResponse struct {
	Symbol    string `json:"symbol"`
	Timeframe string `json:"timeframe"`
	Rate     int    `json:"rate"`
	Status    string `json:"status"`
}
