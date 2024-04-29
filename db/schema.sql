CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE,
  "provider" varchar,
  "refresh_token" int UNIQUE,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "refresh_token" (
  "id" bigserial PRIMARY KEY,
  "token" varchar,
  "user_id" int UNIQUE,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("user_id") REFERENCES "user" ("id")
);


