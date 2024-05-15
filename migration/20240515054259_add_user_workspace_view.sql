-- +goose Up
CREATE VIEW user_workspace_view AS
SELECT u.id, JSON_AGG(DISTINCT w.*) AS workspaces 
from workspace w
LEFT JOIN workspace_user wu ON w.id = wu.workspace_id
LEFT JOIN "user" u ON wu.user_id = u.id
GROUP BY u.id;


-- +goose Down
DROP VIEW user_workspace_view;

