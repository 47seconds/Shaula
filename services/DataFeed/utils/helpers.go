package utils

import (
	"datafeed/models"
)

func SuccessResponse(message string, data any) models.Response {
	return models.Response{
		Code:    200,
		Message: message,
		Status:  true,
		Data:    data,
	}
}

func ErrorResponse(code int, message string) models.Response {
	return models.Response{
		Code:    code,
		Message: message,
		Status:  false,
		Error:   message,
	}
}
