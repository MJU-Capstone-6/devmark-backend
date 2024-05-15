-- +goose Up
CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE,
  "provider" varchar,
  "refresh_token" int UNIQUE,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE "user";

