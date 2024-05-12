-- Create "category" table
CREATE TABLE "public"."category" ("id" bigserial NOT NULL, "name" character varying NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "invite_code" table
CREATE TABLE "public"."invite_code" ("id" bigserial NOT NULL, "workspace_id" integer NULL, "code" character varying NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "invite_code_workspace_id_key" UNIQUE ("workspace_id"));
-- Create "refresh_token" table
CREATE TABLE "public"."refresh_token" ("id" bigserial NOT NULL, "token" character varying NULL, "user_id" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "refresh_token_user_id_key" UNIQUE ("user_id"));
-- Create "user" table
CREATE TABLE "public"."user" ("id" bigserial NOT NULL, "username" character varying NULL, "provider" character varying NULL, "refresh_token" integer NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "user_refresh_token_key" UNIQUE ("refresh_token"), CONSTRAINT "user_username_key" UNIQUE ("username"));
-- Create "workspace" table
CREATE TABLE "public"."workspace" ("id" bigserial NOT NULL, "name" character varying NULL, "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "workspace_category" table
CREATE TABLE "public"."workspace_category" ("workspace_id" bigint NOT NULL, "category_id" bigint NOT NULL, PRIMARY KEY ("workspace_id", "category_id"));
-- Create "workspace_user" table
CREATE TABLE "public"."workspace_user" ("workspace_id" bigint NOT NULL, "user_id" bigint NOT NULL, PRIMARY KEY ("workspace_id", "user_id"));
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
