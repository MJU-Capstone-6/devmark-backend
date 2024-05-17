-- +goose Up
ALTER TABLE "workspace_category" DROP CONSTRAINT "workspace_category_workspace_id_fkey";
ALTER TABLE "workspace_category" ADD CONSTRAINT "workspace_category_workspace_id_fkey"
FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id") ON DELETE CASCADE;

-- +goose Down
ALTER TABLE "workspace_category" DROP CONSTRAINT "workspace_category_workspace_id_fkey";
ALTER TABLE "workspace_category" ADD CONSTRAINT "workspace_category_workspace_id_fkey"
FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id"); 


