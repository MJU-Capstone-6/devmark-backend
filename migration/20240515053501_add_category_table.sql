-- +goose Up
CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE "category";

