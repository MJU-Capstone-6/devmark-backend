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
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);

ALTER TABLE "user" ADD FOREIGN KEY ("refresh_token") REFERENCES "refresh_token" ("id");

CREATE TABLE "workspace" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "invite_code" (
  "id" bigserial PRIMARY KEY,
  "workspace_id" int UNIQUE,
  "code" varchar,
  "expired_at" timestamptz DEFAULT (CURRENT_TIMESTAMP + INTERVAL '1 day'),
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id")
);

CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
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

CREATE VIEW workspace_user_category AS
SELECT w.*, JSON_AGG(DISTINCT c.*) AS categories, JSON_AGG(DISTINCT u.*) AS users
from workspace w
LEFT JOIN workspace_category wc ON w.id = wc.workspace_id
LEFT JOIN category c ON wc.category_id = c.id
LEFT JOIN workspace_user wu ON w.id = wu.workspace_id
LEFT JOIN "user" u ON wu.user_id = u.id
GROUP BY w.id;

CREATE TABLE "bookmark" (
  "id" bigserial PRIMARY KEY,
  "link" varchar,
  "category_id" bigint,
  "workspace_id" bigint,
  "summary" varchar,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("workspace_id") REFERENCES "workspace" ("id"),
  FOREIGN KEY ("category_id") REFERENCES "category" ("id")
);

CREATE TABLE "comment" (
  "id" bigserial PRIMARY KEY,
  "bookmark_id" bigint,
  "user_id" bigint,
  "context" bigint,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY ("bookmark_id") REFERENCES "bookmark" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);

