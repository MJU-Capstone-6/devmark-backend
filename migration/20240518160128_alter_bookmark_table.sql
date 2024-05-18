-- +goose Up
ALTER TABLE "bookmark"
ADD COLUMN "title" varchar;

-- +goose Down
DROP TABLE "bookmark";
