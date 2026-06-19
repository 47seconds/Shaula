import os
import yfinance as yf

stocks = [
    "ADANIENT.NS",
    "ADANIPORTS.NS",
    "JIOFIN.NS",
    "SUZLON.NS",
    "BHEL.NS",
    "IREDA.NS",
    "^NSEBANK",   # Bank Nifty
    "^NSEI"       # Nifty 50
]

os.makedirs("data", exist_ok=True)

for symbol in stocks:
    try:
        df = yf.Ticker(symbol).history(
            period="7d",
            interval="1m"
        )

        filename = (
            f"data/{symbol.replace('.NS', '').replace('^', '')}.csv"
        )

        df.to_csv(filename)

        print(f"Saved {filename} ({len(df)} candles)")
    except Exception as e:
        print(f"Failed {symbol}: {e}")