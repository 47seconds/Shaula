CREATE TABLE candles (
    symbol      TEXT NOT NULL,
    timeframe   TEXT NOT NULL,

    ts          TIMESTAMPTZ NOT NULL,

    open        NUMERIC(18,8) NOT NULL,
    high        NUMERIC(18,8) NOT NULL,
    low         NUMERIC(18,8) NOT NULL,
    close       NUMERIC(18,8) NOT NULL,

    volume      NUMERIC(18,8) NOT NULL,

    PRIMARY KEY (symbol, timeframe, ts)
);

CREATE TABLE strategies (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE
);

CREATE TABLE backtest_runs (
    id                  BIGSERIAL PRIMARY KEY,

    strategy_id         BIGINT NOT NULL
                        REFERENCES strategies(id),

    symbol              TEXT NOT NULL,
    timeframe           TEXT NOT NULL,

    started_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    initial_capital     NUMERIC(18,8) NOT NULL
);

CREATE TABLE ppr_trades (
    id          BIGSERIAL PRIMARY KEY,

    run_id      BIGINT NOT NULL
                REFERENCES backtest_runs(id),

    symbol      TEXT NOT NULL,

    side        TEXT NOT NULL,

    price       NUMERIC(18,8) NOT NULL,

    quantity    NUMERIC(18,8) NOT NULL,

    ts          TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_candles_lookup
ON candles(symbol, timeframe, ts);

CREATE INDEX idx_trades_run
ON ppr_trades(run_id);