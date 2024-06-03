-- +goose Up
CREATE VIEW top_workspace_categories AS
WITH bookmark_counts AS (
    SELECT
        workspace_id,
        category_id,
        COUNT(*) AS bookmark_count
    FROM
        bookmark
    GROUP BY
        workspace_id, category_id
),
top_categories AS (
    SELECT
        workspace_id,
        category_id,
        bookmark_count,
        ROW_NUMBER() OVER (PARTITION BY workspace_id ORDER BY bookmark_count DESC) AS rank
    FROM
        bookmark_counts
)
SELECT
    workspace_id,
    category_id,
    bookmark_count
FROM
    top_categories
WHERE
    rank <= 3;
-- +goose Down
DROP VIEW top_3_categories;
