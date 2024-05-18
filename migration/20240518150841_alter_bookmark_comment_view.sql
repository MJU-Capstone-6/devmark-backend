-- +goose Up
DROP VIEW bookmark_comment;

-- +goose Down
CREATE VIEW bookmark_comment AS
SELECT b.*, JSON_AGG(DISTINCT c.*) AS comments
from bookmark b
LEFT JOIN "comment" c ON c.bookmark_id = b.id
GROUP BY b.id;
