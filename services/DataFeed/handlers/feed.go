type Feed interface {
    Next() (*models.Candle, error)
}

type PGFeed struct {
    rows pgx.Rows
}