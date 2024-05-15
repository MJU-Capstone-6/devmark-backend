-- +goose Up
CREATE TABLE "workspace" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "description" varchar,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE "workspace";

