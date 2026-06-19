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
	Speed     int    `json:"speed"`
	Status    string `json:"status"`
}
