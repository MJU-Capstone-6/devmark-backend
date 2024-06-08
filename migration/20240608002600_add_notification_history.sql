-- +goose Up
CREATE TABLE notification_history (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "notification_title" varchar,
  "is_read" boolean DEFAULT false,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);

CREATE VIEW unread_notifications AS
SELECT 
"user".id, "user".username, JSON_AGG(DISTINCT nh.*) AS notifications
FROM 
  notification_history nh
JOIN "user" ON "user".id = nh.user_id
WHERE nh.is_read = false
GROUP BY "user".id;

-- +goose Down
DROP VIEW unread_notifications;
DROP TABLE notification_history;
