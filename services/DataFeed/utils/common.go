package utils

const (
	OK                  = 200
	BAD_REQUEST         = 400
	UNAUTHORIZED        = 401
	FORBIDDEN           = 403
	NOT_FOUND           = 404
	INTERNAL_SERVER_ERR = 500
)

const DB_QUERY_ALL = `
SELECT
	symbol,
	timeframe,
	ts,
	open,
	high,
	low,
	close,
	volume
FROM candles
WHERE symbol = $1
AND timeframe = $2
ORDER BY ts ASC
`

const DB_QUERY_LIMIT = `
SELECT
    symbol,
    timeframe,
    ts,
    open,
    high,
    low,
    close,
    volume
FROM candles
WHERE symbol = $1
AND timeframe = $2
ORDER BY ts ASC
LIMIT $3
`