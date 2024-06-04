-- +goose Up
ALTER TABLE recommend_link
DROP CONSTRAINT "recommend_link_workspace_id_fkey";

ALTER TABLE recommend_link
ADD CONSTRAINT "fk_workspace"
FOREIGN KEY (workspace_id) REFERENCES "workspace" ("id");

-- +goose Down
DROP TABLE recommend_link;
