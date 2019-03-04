-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS acts_artistes(id INTEGER PRIMARY KEY AUTOINCREMENT,
  act_id INTEGER NOT NULL,
  artiset_id INTEGER NOT NULL,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS acts_artistes;
