-- Modify "category" table
ALTER TABLE "public"."category" ALTER COLUMN "created_at" SET NOT NULL, ALTER COLUMN "updated_at" SET NOT NULL;
-- Modify "invite_code" table
ALTER TABLE "public"."invite_code" ALTER COLUMN "created_at" SET NOT NULL, ALTER COLUMN "updated_at" SET NOT NULL;
-- Modify "refresh_token" table
ALTER TABLE "public"."refresh_token" ALTER COLUMN "created_at" SET NOT NULL, ALTER COLUMN "updated_at" SET NOT NULL;
-- Modify "user" table
ALTER TABLE "public"."user" ALTER COLUMN "created_at" SET NOT NULL, ALTER COLUMN "updated_at" SET NOT NULL;
-- Modify "workspace" table
ALTER TABLE "public"."workspace" ALTER COLUMN "created_at" SET NOT NULL, ALTER COLUMN "updated_at" SET NOT NULL;
