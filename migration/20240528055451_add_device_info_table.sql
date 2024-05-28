-- +goose Up
CREATE TABLE "device_info" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "agent_header" VARCHAR NOT NULL,
  "device_id" VARCHAR,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);

-- +goose Down
DROP TABLE "device_info";

