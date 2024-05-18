-- +goose Up
CREATE VIEW bookmark_comment AS
SELECT 
    b.*, 
    JSON_AGG(
        DISTINCT JSONB_BUILD_OBJECT(
            'comment_id', c.id,
            'comment_context', c.comment_context,
            'user_id', u.id,
            'username', u.username
        )
    ) AS comments
FROM 
    bookmark b
LEFT JOIN 
    "comment" c ON c.bookmark_id = b.id
LEFT JOIN 
    "user" u ON c.user_id = u.id
GROUP BY 
    b.id;
-- +goose Down
DROP VIEW bookmark_comment;
