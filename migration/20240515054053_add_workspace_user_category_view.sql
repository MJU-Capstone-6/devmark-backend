-- +goose Up
CREATE VIEW workspace_user_category AS
SELECT w.*, JSON_AGG(DISTINCT c.*) AS categories, JSON_AGG(DISTINCT u.*) AS users
from workspace w
LEFT JOIN workspace_category wc ON w.id = wc.workspace_id
LEFT JOIN category c ON wc.category_id = c.id
LEFT JOIN workspace_user wu ON w.id = wu.workspace_id
LEFT JOIN "user" u ON wu.user_id = u.id
GROUP BY w.id;


-- +goose Down
DROP VIEW workspace_user_category;

