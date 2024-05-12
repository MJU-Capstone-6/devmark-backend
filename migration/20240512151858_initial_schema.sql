-- Modify "category" table
ALTER TABLE "public"."category" DROP COLUMN "created_at", DROP COLUMN "updated_at";
-- Modify "user" table
ALTER TABLE "public"."user" DROP COLUMN "created_at", DROP COLUMN "updated_at";
