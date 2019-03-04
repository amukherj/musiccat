-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE acts_artistes
RENAME COLUMN artiset_id TO artiste_id;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE acts_artistes
RENAME COLUMN artiste_id TO artiset_id;
