-- +goose Up
ALTER TABLE "refresh_token" DROP CONSTRAINT "refresh_token_user_id_fkey";
ALTER TABLE "refresh_token" ADD CONSTRAINT "refresh_token_user_id_fkey"
FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "invite_code" DROP CONSTRAINT "invite_code_workspace_id_fkey";
ALTER TABLE "invite_code" ADD CONSTRAINT "invite_code_workspace_id_fkey"
FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id") ON DELETE CASCADE;

ALTER TABLE "bookmark" DROP CONSTRAINT "bookmark_workspace_id_fkey";
ALTER TABLE "bookmark" ADD CONSTRAINT "bookmark_workspace_id_fkey"
FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id") ON DELETE CASCADE;

-- +goose Down
ALTER TABLE "refresh_token" DROP CONSTRAINT "refresh_token_user_id_fkey";
ALTER TABLE "refresh_token" ADD CONSTRAINT "refresh_token_user_id_fkey"
FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "invite_code" DROP CONSTRAINT "invite_code_workspace_id_fkey";
ALTER TABLE "invite_code" ADD CONSTRAINT "invite_code_workspace_id_fkey"
FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id");

ALTER TABLE "bookmark" DROP CONSTRAINT "bookmark_workspace_id_fkey";
ALTER TABLE "bookmark" ADD CONSTRAINT "bookmark_workspace_id_fkey"
FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id");


