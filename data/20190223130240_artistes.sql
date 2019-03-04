-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS artistes(id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  artiste_type TEXT NOT NULL,
  start_year INTEGER,
  last_active_year INTEGER,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS artistes;
