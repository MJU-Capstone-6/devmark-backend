-- +goose Up
CREATE TABLE "bookmark" (
  "id" bigserial PRIMARY KEY,
  "link" varchar,
  "category_id" bigint,
  "workspace_id" bigint,
  "summary" varchar NULL,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id"),
  FOREIGN KEY ("category_id") REFERENCES "category" ("id")
);

-- +goose StatementBegin
ALTER TABLE "workspace" ADD COLUMN "bookmark_count" integer DEFAULT 0;

CREATE OR REPLACE FUNCTION update_workspace_bookmark_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE "workspace"
    SET "bookmark_count" = (
        SELECT COUNT(*)
        FROM "bookmark"
        WHERE "workspace_id" = NEW."workspace_id"
    )
    WHERE "id" = NEW."workspace_id";

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_workspace_bookmark_count
AFTER INSERT OR UPDATE OR DELETE ON "bookmark"
FOR EACH ROW EXECUTE FUNCTION update_workspace_bookmark_count();

ALTER TABLE "workspace" ADD COLUMN "user_count" integer DEFAULT 0;

CREATE OR REPLACE FUNCTION update_workspace_user_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE "workspace" SET "user_count" = (
      SELECT COUNT("user_id")
      FROM "workspace_user"
      WHERE "workspace_id" = "workspace"."id"
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_workspace_user_count
AFTER INSERT OR UPDATE OR DELETE ON "workspace_user"
FOR EACH ROW EXECUTE FUNCTION update_workspace_user_count();
-- +goose StatementEnd

-- +goose Down
DROP TABLE "bookmark";

