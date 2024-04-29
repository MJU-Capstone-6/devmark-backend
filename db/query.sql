-- name: FindUserByUsername :one
SELECT * FROM "user" WHERE "username" = $1 LIMIT 1;

-- name: Create :one
INSERT INTO "user" (
  "username",
  "provider",
  "refresh_token"
) VALUES (
  $1,
  $2,
  $3
) RETURNING *;
