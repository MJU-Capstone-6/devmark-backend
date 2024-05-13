-- Modify "category" table
ALTER TABLE "public"."category" ADD CONSTRAINT "category_name_key" UNIQUE ("name");
