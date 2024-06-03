-- +goose Up
ALTER TABLE bookmark 
ADD COLUMN is_read boolean DEFAULT false;

-- +goose Down
ALTER TABLE bookmark
DROP COLUMN is_read;

