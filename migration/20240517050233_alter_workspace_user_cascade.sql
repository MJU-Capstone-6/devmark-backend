-- +goose Up
ALTER TABLE "workspace_user" DROP CONSTRAINT "workspace_user_workspace_id_fkey";
ALTER TABLE "workspace_user" ADD CONSTRAINT "workspace_user_workspace_id_fkey"
FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id") ON DELETE CASCADE;

-- +goose Down
ALTER TABLE "workspace_user" DROP CONSTRAINT "workspace_user_workspace_id_fkey";
ALTER TABLE "workspace_user" ADD CONSTRAINT "workspace_user_workspace_id_fkey"
FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id") ON DELETE CASCADE;

