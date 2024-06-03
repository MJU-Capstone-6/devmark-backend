-- +goose Up
ALTER TABLE "recommend_link" DROP CONSTRAINT IF EXISTS "fk_user_id";
ALTER TABLE "recommend_link" DROP COLUMN IF EXISTS "user_id";

ALTER TABLE "recommend_link" ADD COLUMN "category_id" bigint;
ALTER TABLE "recommend_link" ADD CONSTRAINT "fk_category_id" FOREIGN KEY ("category_id") REFERENCES "category" ("id");

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
    top_categories.workspace_id,
    top_categories.category_id,
    bookmark_count,
    JSON_AGG(DISTINCT r.*) AS recommend_links
FROM
    top_categories
JOIN recommend_link r ON r.workspace_id = top_categories.workspace_id AND r.category_id = top_categories.category_id
WHERE
    rank <= 3
GROUP BY top_categories.workspace_id, top_categories.category_id, top_categories.bookmark_count;
-- +goose Down
DROP TABLE top_workspace_categories;
