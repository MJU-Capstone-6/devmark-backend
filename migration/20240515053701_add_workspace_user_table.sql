-- +goose Up
CREATE TABLE "workspace_user" (
  "workspace_id" bigint,
  "user_id" bigint,
  FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "user" ("id"),
  PRIMARY KEY ("workspace_id", "user_id")
);


-- +goose Down
DROP TABLE "workspace_user";

