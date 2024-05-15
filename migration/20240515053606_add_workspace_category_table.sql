-- +goose Up
CREATE TABLE "workspace_category" (
  "workspace_id" bigint,
  "category_id" bigint,
  FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id"),
  FOREIGN KEY ("category_id") REFERENCES "category" ("id"),
  PRIMARY KEY ("workspace_id", "category_id")
);


-- +goose Down
DROP TABLE "workspace_category";

