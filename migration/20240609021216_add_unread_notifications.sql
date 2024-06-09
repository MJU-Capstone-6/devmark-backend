-- +goose Up
CREATE VIEW unread_notifications AS
SELECT 
"user".id AS user_id, "user".username, JSON_AGG(DISTINCT nh.*) AS notifications, bookmark.id AS bookmark_id, bookmark.link AS bookmark_link, bookmark.summary AS bookmark_summary, bookmark.title AS bookmark_title
FROM 
  notification_history nh
JOIN "user" ON "user".id = nh.user_id
JOIN bookmark ON bookmark.id = nh.bookmark_id
WHERE nh.is_read = false
GROUP BY "user".id, bookmark.id;


-- +goose Down
DROP VIEW unread_notifications;

