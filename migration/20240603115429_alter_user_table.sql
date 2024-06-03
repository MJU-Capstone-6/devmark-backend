-- +goose Up
DROP TABLE device_info;

-- +goose Down
CREATE OR REPLACE TABLE "device_info" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "registration_token" varchar(500),
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);
