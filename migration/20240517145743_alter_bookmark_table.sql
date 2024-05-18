-- +goose Up
ALTER TABLE "bookmark"
ADD COLUMN "user_id" bigint;

ALTER TABLE "bookmark"
ADD CONSTRAINT "fk_user"
FOREIGN KEY ("user_id") REFERENCES "user" ("id");

-- +goose Down
ALTER TABLE "bookmark"
DROP CONSTRAINT "fk_user";

ALTER TABLE "bookmark"
DROP COLUMN "user_id";
