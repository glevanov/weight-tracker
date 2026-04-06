CREATE TABLE IF NOT EXISTS users (
  username TEXT PRIMARY KEY,
  password TEXT NOT NULL,
  salt TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS weights (
  id BIGSERIAL PRIMARY KEY,
  username TEXT NOT NULL,
  weight DOUBLE PRECISION NOT NULL,
  timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_weights_username_timestamp
  ON weights (username, timestamp);
