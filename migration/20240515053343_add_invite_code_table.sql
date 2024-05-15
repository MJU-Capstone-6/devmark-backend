-- +goose Up
create table "invite_code" (
  "id" bigserial primary key,
  "workspace_id" int unique,
  "code" varchar unique,
  "expired_at" timestamptz default (current_timestamp + interval '1 day'),
  "created_at" timestamptz default current_timestamp not null,
  "updated_at" timestamptz default current_timestamp not null,
  foreign key ("workspace_id") references "workspace" ("id")
);


-- +goose Down
DROP TABLE "invite_code";

