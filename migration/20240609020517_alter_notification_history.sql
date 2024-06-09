-- +goose Up
ALTER TABLE notification_history
ADD COLUMN bookmark_id bigint;

ALTER TABLE notification_history
ADD CONSTRAINT fk_bookmark
FOREIGN KEY (bookmark_id) REFERENCES bookmark(id);
-- +goose Down
ALTER TABLE notification_history
DROP COLUMN bookmark_id bigint;

ALTER TABLE notification_history
DROP CONSTRAINT fk_bookmark
FOREIGN KEY (bookmark_id) REFERENCES bookmark(id);

