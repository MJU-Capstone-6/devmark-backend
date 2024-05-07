-- name: FindUserByUsername :one
SELECT * FROM "user" WHERE "username" = $1 LIMIT 1;

-- name: CreateRefreshToken :one
INSERT INTO refresh_token (token, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: FindRefreshTokenByUserID :one
SELECT * from "refresh_token" WHERE "user_id" = $1 LIMIT 1;

-- name: UpdateRefreshToken :one
UPDATE refresh_token
SET
    token = COALESCE(NULLIF($2, ''), token),
    user_id = COALESCE(NULLIF($3, ''), user_id),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *; 

