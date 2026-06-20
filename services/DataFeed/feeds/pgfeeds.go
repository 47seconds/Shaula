package feeds

import (
	"io"

	"datafeed/models"

	"github.com/jackc/pgx/v4"
)

type PGFeed struct {
	rows pgx.Rows
}

func NewPGFeed(rows pgx.Rows) *PGFeed {
	return &PGFeed{
		rows: rows,
	}
}

func (f *PGFeed) Close() {
	f.rows.Close()
}

func (f *PGFeed) Next() (*models.Candle, error) {
    if !f.rows.Next() {
        if err := f.rows.Err(); err != nil {
            return nil, err
        }

        return nil, io.EOF
    }

	var c models.Candle

	err := f.rows.Scan(
		&c.Symbol,
		&c.Timeframe,
		&c.Timestamp,
		&c.Open,
		&c.High,
		&c.Low,
		&c.Close,
		&c.Volume,
	)

	if err != nil {
		return nil, err
	}

	return &c, nil
}