-- +goose Up
CREATE TABLE "recommend_link" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "workspace_id" bigint,
  "link" varchar,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("user_id") REFERENCES "user" ("id"),
  FOREIGN KEY ("workspace_id") REFERENCES "user" ("id")
);

-- +goose Down
DROP TABLE "recommend_link";
