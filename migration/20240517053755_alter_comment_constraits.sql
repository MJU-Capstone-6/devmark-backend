-- +goose Up
ALTER TABLE "comment" DROP CONSTRAINT "comment_user_id_fkey";
ALTER TABLE "comment" ADD CONSTRAINT "comment_user_id_fkey"
FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "comment" DROP CONSTRAINT "comment_bookmark_id_fkey";
ALTER TABLE "comment" ADD CONSTRAINT "comment_bookmark_id_fkey"
FOREIGN KEY ("bookmark_id") REFERENCES "bookmark" ("id") ON DELETE CASCADE;

-- +goose Down
ALTER TABLE "comment" DROP CONSTRAINT "comment_user_id_fkey";
ALTER TABLE "comment" ADD CONSTRAINT "comment_user_id_fkey"
FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "comment" DROP CONSTRAINT "comment_bookmark_id_fkey";
ALTER TABLE "comment" ADD CONSTRAINT "comment_bookmark_id_fkey"
FOREIGN KEY ("bookmark_id") REFERENCES "bookmark" ("id") ON DELETE CASCADE;
