-- +goose Up
CREATE OR REPLACE VIEW workspace_user_category AS
SELECT 
    w.*, 
    JSON_AGG(DISTINCT c.*) AS categories, 
    JSON_AGG(DISTINCT JSONB_BUILD_OBJECT('id', u.id, 'username', u.username,'provider', u.provider, 'created_at', u.created_at, 'updated_at', u.updated_at)) AS users
FROM 
    workspace w
LEFT JOIN 
    workspace_category wc ON w.id = wc.workspace_id
LEFT JOIN 
    category c ON wc.category_id = c.id
LEFT JOIN 
    workspace_user wu ON w.id = wu.workspace_id
LEFT JOIN 
    "user" u ON wu.user_id = u.id
GROUP BY 
    w.id;

-- +goose Down
DROP VIEW workspace_user_category;

