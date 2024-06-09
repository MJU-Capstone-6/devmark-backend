-- +goose Up
DROP VIEW unread_notifications;
CREATE VIEW unread_notifications AS
SELECT 
    nh.id AS notification_id,
    nh.user_id,
    nh.notification_title,
    nh.is_read,
    nh.created_at,
    nh.updated_at,
    nh.bookmark_id,
    nh.workspace_id,
    ub.bookmarks,
    ub.device_infos
FROM 
    notification_history nh
LEFT JOIN 
    unread_bookmark ub ON nh.user_id = ub.user_id AND nh.workspace_id = ub.workspace_id;
-- +goose Down
DROP VIEW unread_notifications;
