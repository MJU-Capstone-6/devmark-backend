-- +goose Up
CREATE VIEW unread_bookmark AS
SELECT user_id, workspace_id, JSON_AGG(DISTINCT b.*) AS bookmarks
FROM bookmark b
WHERE is_read = false
GROUP BY user_id, workspace_id;

-- +goose Down
DROP VIEW unread_bookmark;
