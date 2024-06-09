-- +goose Up
CREATE VIEW unread_notifications AS
SELECT 
"user".id AS user_id, "user".username, JSON_AGG(DISTINCT nh.*) AS notifications 
FROM 
  notification_history nh
JOIN "user" ON "user".id = nh.user_id
WHERE nh.is_read = false
GROUP BY "user".id;

-- +goose Down
DROP VIEW unread_notifications;

