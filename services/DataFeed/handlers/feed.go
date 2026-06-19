package handlers

import (
    "github.com/jackc/pgx/v4"
    "datafeed/models"
)

type Feed interface {
    Next() (*models.Candle, error)
}

type PGFeed struct {
    rows pgx.Rows
}