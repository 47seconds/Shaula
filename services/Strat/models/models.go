package models

import "time"

type Candle struct {
	Symbol    string    `json:"symbol"`
	Timeframe string    `json:"timeframe"`
	Timestamp time.Time `json:"timestamp"`

	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`

	Volume float64 `json:"volume"`
}

type CandleMessage struct {
	Type string        `json:"type"`
	Data Candle `json:"data"`
}