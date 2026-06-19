package utils

import (
	"os"
	"strconv"

	"strat/models"
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

func StringToInt(s string, defaultValue int) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        return defaultValue
    }
    return n
}

func GetEnvString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return StringToInt(value, defaultValue)
}
