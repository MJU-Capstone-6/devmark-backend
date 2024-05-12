-- Modify "category" table
ALTER TABLE "public"."category" ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz;
-- Modify "user" table
ALTER TABLE "public"."user" ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz;
