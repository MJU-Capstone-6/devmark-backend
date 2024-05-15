-- +goose Up
CREATE TABLE "comment" (
  "id" bigserial PRIMARY KEY,
  "bookmark_id" bigint,
  "user_id" bigint,
  "context" bigint,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("bookmark_id") REFERENCES "bookmark" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);


-- +goose Down
DROP TABLE "comment";

