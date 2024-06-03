-- +goose Up
CREATE OR REPLACE VIEW unread_bookmark AS
SELECT b.user_id, workspace_id, JSON_AGG(DISTINCT b.*) AS bookmarks, JSON_AGG(DISTINCT d.*) AS device_infos
FROM bookmark b
JOIN device_info d ON b.user_id = d.user_id
WHERE is_read = false
GROUP BY b.user_id, workspace_id;

-- +goose Down
DROP VIEW unread_bookmark;
