package models

import "time"

type SignalType string

const (
	Hold SignalType = "HOLD"
	Buy  SignalType = "BUY"
	Sell SignalType = "SELL"
)

type Signal struct {
	Symbol    string     `json:"symbol"`
	Timeframe string     `json:"timeframe"`
	Type      SignalType `json:"type"`
	Price     float64    `json:"price"`
	Timestamp time.Time  `json:"timestamp"`
}
