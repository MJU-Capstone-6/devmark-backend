-- +goose Up
CREATE TABLE "workspace_code" (
  "id" bigserial PRIMARY KEY,
  "workspace_id" bigint,
  "code" varchar UNIQUE,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id")
);

-- +goose Down
DROP TABLE "workspace_code";
