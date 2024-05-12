-- Modify "category" table
ALTER TABLE "public"."category" ADD COLUMN "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, ADD COLUMN "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP;
-- Modify "user" table
ALTER TABLE "public"."user" ADD COLUMN "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP, ADD COLUMN "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP;
