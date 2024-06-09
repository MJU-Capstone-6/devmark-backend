-- +goose Up
DROP VIEW unread_notifications;

-- +goose Down
CREATE VIEW unread_notifications AS
SELECT 
"user".id, "user".username, JSON_AGG(DISTINCT nh.*) AS notifications
FROM 
  notification_history nh
JOIN "user" ON "user".id = nh.user_id
WHERE nh.is_read = false
GROUP BY "user".id;



