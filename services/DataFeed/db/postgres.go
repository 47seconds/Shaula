package db

import (
	"context"
	"log"
	"fmt"

	"datafeed/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() error {
	log.Println("Connecting to PostgreSQL...")
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		utils.GetEnvString("POSTGRES_HOST", "localhost"),
		utils.GetEnvString("POSTGRES_PORT", "5432"),
		utils.GetEnvString("POSTGRES_USER", "postgres"),
		utils.GetEnvString("POSTGRES_PASSWORD", ""),
		utils.GetEnvString("POSTGRES_DB", "postgres"),
	)

	pool, err := pgxpool.Connect(
		context.Background(),
		dsn,
	)
	if err != nil {
		return err
	}

	Pool = pool

	log.Println("Connected to PostgreSQL")

	return nil
}