-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS songs(id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  artiste_id INTEGER NOT NULL,
  album_id INTEGER,
  genre_id INTEGER NOT NULL,
  credits TEXT,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL);



-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE songs;
