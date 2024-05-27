-- +goose Up
CREATE VIEW workspace_user_bookmark_count AS
WITH user_bookmark_counts AS (
    SELECT 
        wu.workspace_id,
        wu.user_id,
        u.username,
        COUNT(b.id) AS bookmark_count
    FROM 
        workspace_user wu
    LEFT JOIN 
        bookmark b ON wu.user_id = b.user_id AND wu.workspace_id = b.workspace_id
    LEFT JOIN
        "user" u ON wu.user_id = u.id
    GROUP BY 
        wu.workspace_id, wu.user_id, u.username
)
SELECT 
    workspace_id,
    json_agg(
        json_build_object(
            'user_id', user_id,
            'username', username,
            'bookmark_count', bookmark_count
        )
    ) AS user_bookmark_count
FROM 
    user_bookmark_counts
GROUP BY 
    workspace_id;
-- +goose Down
DROP VIEW "workspace_user_bookmark_count";
