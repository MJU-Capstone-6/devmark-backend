-- +goose Up
CREATE VIEW bookmark_comment AS
SELECT b.*, JSON_AGG(DISTINCT c.*) AS comments
from bookmark b
LEFT JOIN "comment" c ON c.bookmark_id = b.id
GROUP BY b.id;

-- +goose Down
DROP VIEW bookmark_comment;
