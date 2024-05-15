-- +goose Up
ALTER TABLE "comment"
RENAME COLUMN "context" TO "comment_context";
-- +goose Down
DROP TABLE "comment";
