import os
import pandas as pd
import psycopg2

conn = psycopg2.connect(
    host="localhost",
    port=5432,
    dbname="shaula",
    user="shaula",
    password=""
)

cur = conn.cursor()

DATA_DIR = "./data"

for filename in os.listdir(DATA_DIR):
    if not filename.endswith(".csv"):
        continue

    symbol = filename.replace(".csv", "")
    path = os.path.join(DATA_DIR, filename)

    print(f"Importing {symbol}")

    df = pd.read_csv(path)

    for _, row in df.iterrows():
        cur.execute(
            """
            INSERT INTO candles (
                symbol,
                timeframe,
                ts,
                open,
                high,
                low,
                close,
                volume
            )
            VALUES (
                %s,%s,%s,%s,%s,%s,%s,%s
            )
            ON CONFLICT DO NOTHING
            """,
            (
                symbol,
                "1m",
                row["Datetime"],
                float(row["Open"]),
                float(row["High"]),
                float(row["Low"]),
                float(row["Close"]),
                float(row["Volume"])
            )
        )

    conn.commit()

cur.close()
conn.close()

print("Done")