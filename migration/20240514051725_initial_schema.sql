-- Create "bookmark" table
CREATE TABLE "public"."bookmark" ("id" bigserial NOT NULL, "link" character varying NULL, "category_id" bigint NULL, "workspace_id" bigint NULL, "summary" character varying NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "category" table
CREATE TABLE "public"."category" ("id" bigserial NOT NULL, "name" character varying NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "category_name_key" UNIQUE ("name"));
-- Create "comment" table
CREATE TABLE "public"."comment" ("id" bigserial NOT NULL, "bookmark_id" bigint NULL, "user_id" bigint NULL, "context" bigint NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "invite_code" table
CREATE TABLE "public"."invite_code" ("id" bigserial NOT NULL, "workspace_id" integer NULL, "code" character varying NULL, "expired_at" timestamptz NULL DEFAULT (CURRENT_TIMESTAMP + '1 day'::interval), "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "invite_code_workspace_id_key" UNIQUE ("workspace_id"));
-- Create "refresh_token" table
CREATE TABLE "public"."refresh_token" ("id" bigserial NOT NULL, "token" character varying NULL, "user_id" integer NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "refresh_token_user_id_key" UNIQUE ("user_id"));
-- Create "user" table
CREATE TABLE "public"."user" ("id" bigserial NOT NULL, "username" character varying NULL, "provider" character varying NULL, "refresh_token" integer NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "user_refresh_token_key" UNIQUE ("refresh_token"), CONSTRAINT "user_username_key" UNIQUE ("username"));
-- Create "workspace" table
CREATE TABLE "public"."workspace" ("id" bigserial NOT NULL, "name" character varying NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "workspace_category" table
CREATE TABLE "public"."workspace_category" ("workspace_id" bigint NOT NULL, "category_id" bigint NOT NULL, PRIMARY KEY ("workspace_id", "category_id"));
-- Create "workspace_user" table
CREATE TABLE "public"."workspace_user" ("workspace_id" bigint NOT NULL, "user_id" bigint NOT NULL, PRIMARY KEY ("workspace_id", "user_id"));
-- Modify "bookmark" table
ALTER TABLE "public"."bookmark" ADD CONSTRAINT "bookmark_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "public"."category" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "bookmark_workspace_id_fkey" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspace" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "comment" table
ALTER TABLE "public"."comment" ADD CONSTRAINT "comment_bookmark_id_fkey" FOREIGN KEY ("bookmark_id") REFERENCES "public"."bookmark" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "comment_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "invite_code" table
ALTER TABLE "public"."invite_code" ADD CONSTRAINT "invite_code_workspace_id_fkey" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspace" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "refresh_token" table
ALTER TABLE "public"."refresh_token" ADD CONSTRAINT "refresh_token_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "user" table
ALTER TABLE "public"."user" ADD CONSTRAINT "user_refresh_token_fkey" FOREIGN KEY ("refresh_token") REFERENCES "public"."refresh_token" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "workspace_category" table
ALTER TABLE "public"."workspace_category" ADD CONSTRAINT "workspace_category_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "public"."category" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "workspace_category_workspace_id_fkey" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspace" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "workspace_user" table
ALTER TABLE "public"."workspace_user" ADD CONSTRAINT "workspace_user_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD CONSTRAINT "workspace_user_workspace_id_fkey" FOREIGN KEY ("workspace_id") REFERENCES "public"."workspace" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Create "user_workspace_view" view
CREATE VIEW "public"."user_workspace_view" ("id", "workspaces") AS SELECT u.id,
    json_agg(DISTINCT w.*) AS workspaces
   FROM ((workspace w
     LEFT JOIN workspace_user wu ON ((w.id = wu.workspace_id)))
     LEFT JOIN "user" u ON ((wu.user_id = u.id)))
  GROUP BY u.id;
-- Create "workspace_user_category" view
CREATE VIEW "public"."workspace_user_category" ("id", "name", "created_at", "updated_at", "categories", "users") AS SELECT w.id,
    w.name,
    w.created_at,
    w.updated_at,
    json_agg(DISTINCT c.*) AS categories,
    json_agg(DISTINCT u.*) AS users
   FROM ((((workspace w
     LEFT JOIN workspace_category wc ON ((w.id = wc.workspace_id)))
     LEFT JOIN category c ON ((wc.category_id = c.id)))
     LEFT JOIN workspace_user wu ON ((w.id = wu.workspace_id)))
     LEFT JOIN "user" u ON ((wu.user_id = u.id)))
  GROUP BY w.id;
