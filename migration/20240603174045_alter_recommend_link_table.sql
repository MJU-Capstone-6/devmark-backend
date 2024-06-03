-- +goose Up
ALTER TABLE "recommend_link" ADD COLUMN "title" varchar;
-- +goose Down
ALTER TABLE "recommend_link" DROP COLUMN "title" varchar;
