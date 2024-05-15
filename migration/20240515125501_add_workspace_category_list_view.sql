-- +goose Up
CREATE VIEW workspace_category_list AS
SELECT w.*, JSON_AGG(DISTINCT c.*) AS categories
FROM workspace w
LEFT JOIN workspace_category wc ON wc.workspace_id = w.id
LEFT JOIN category c ON wc.category_id = c.id
GROUP BY w.id;

-- +goose Down
DROP VIEW workspace_category_list;
