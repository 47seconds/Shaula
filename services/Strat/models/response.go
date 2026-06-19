package models

type Response struct {
	Code	int    `json:"code"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data	any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
