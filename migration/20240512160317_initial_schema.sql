-- Modify "invite_code" table
ALTER TABLE "public"."invite_code" ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz, ADD COLUMN "expired_at" timestamptz NULL DEFAULT (CURRENT_TIMESTAMP + '1 day'::interval);
