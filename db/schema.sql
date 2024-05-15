CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE,
  "provider" varchar,
  "refresh_token" int UNIQUE,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "refresh_token" (
  "id" bigserial PRIMARY KEY,
  "token" varchar,
  "user_id" int UNIQUE,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);

ALTER TABLE "user" ADD FOREIGN KEY ("refresh_token") REFERENCES "refresh_token" ("id");

CREATE TABLE "workspace" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "description" varchar,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "invite_code" (
  "id" bigserial PRIMARY KEY,
  "workspace_id" int UNIQUE,
  "code" varchar UNIQUE,
  "expired_at" timestamptz DEFAULT (CURRENT_TIMESTAMP + INTERVAL '1 day'),
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id")
);

CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "workspace_category" (
  "workspace_id" bigint,
  "category_id" bigint,
  FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id"),
  FOREIGN KEY ("category_id") REFERENCES "category" ("id"),
  PRIMARY KEY ("workspace_id", "category_id")
);

CREATE TABLE "workspace_user" (
  "workspace_id" bigint,
  "user_id" bigint,
  FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "user" ("id"),
  PRIMARY KEY ("workspace_id", "user_id")
);

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


CREATE VIEW workspace_user_category AS
SELECT w.*, JSON_AGG(DISTINCT c.*) AS categories, JSON_AGG(DISTINCT u.*) AS users
from workspace w
LEFT JOIN workspace_category wc ON w.id = wc.workspace_id
LEFT JOIN category c ON wc.category_id = c.id
LEFT JOIN workspace_user wu ON w.id = wu.workspace_id
LEFT JOIN "user" u ON wu.user_id = u.id
GROUP BY w.id;


CREATE TABLE "comment" (
  "id" bigserial PRIMARY KEY,
  "bookmark_id" bigint,
  "user_id" bigint,
  "context" text,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("bookmark_id") REFERENCES "bookmark" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);

CREATE VIEW user_workspace_view AS
SELECT u.id, JSON_AGG(DISTINCT w.*) AS workspaces 
from workspace w
LEFT JOIN workspace_user wu ON w.id = wu.workspace_id
LEFT JOIN "user" u ON wu.user_id = u.id
GROUP BY u.id;

CREATE VIEW bookmark_comment AS
SELECT b.*, JSON_AGG(DISTINCT c.*) AS comments, 
FROM bookmark b
LEFT JOIN "comment" c ON c.bookmark_id = bookmark.id
GROUP BY b.id;

CREATE VIEW workspace_category_list AS
SELECT w.*, JSON_AGG(DISTINCT c.*) AS categories, 
FROM workspace w
LEFT JOIN workspace_category wc ON wc.workspace_id = w.id
LEFT JOIN category c ON wc.category_id = c.id
GROUP BY w.id;
