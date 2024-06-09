-- +goose Up
ALTER TABLE notification_history
ADD COLUMN workspace_id bigint;

ALTER TABLE notification_history
ADD CONSTRAINT fk_workspace
FOREIGN KEY (workspace_id) REFERENCES workspace(id);

-- +goose Down
ALTER TABLE notification_history
DROP COLUMN workspace_id bigint;

ALTER TABLE notification_history
DROP CONSTRAINT fk_workspace
FOREIGN KEY (workspace_id) REFERENCES workspace(id);


