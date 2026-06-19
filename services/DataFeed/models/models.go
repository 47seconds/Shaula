type Candle struct {
    Symbol    string
    Timeframe string

    Timestamp time.Time

    Open   float64
    High   float64
    Low    float64
    Close  float64

    Volume float64
}

type Feed interface {
    Next() (*Candle, error)
}

type PGFeed struct {
    rows pgx.Rows
}