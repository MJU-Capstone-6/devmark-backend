-- +goose Up
CREATE TABLE "refresh_token" (
  "id" bigserial PRIMARY KEY,
  "token" varchar,
  "user_id" int UNIQUE,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);

ALTER TABLE "user" ADD FOREIGN KEY ("refresh_token") REFERENCES "refresh_token" ("id");

-- +goose Down
DROP TABLE "refresh_token";

