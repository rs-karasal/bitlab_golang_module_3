-- +goose Up
ALTER TABLE products ADD COLUMN description TEXT;

-- +goose Down
ALTER TABLE products DROP COLUMN IF EXISTS description;
