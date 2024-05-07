-- name: FindUserByUsername :one
SELECT * FROM "user" WHERE "username" = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "user" (username, provider)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUser :one
Update "user"
SET 
    refresh_token = $1,
    username = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $3
RETURNING *;

-- name: CreateRefreshToken :one
INSERT INTO refresh_token (token, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: FindRefreshTokenByUserID :one
SELECT * from "refresh_token" WHERE "user_id" = $1 LIMIT 1;

-- name: UpdateRefreshToken :one
UPDATE refresh_token
SET
    token = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $2 
RETURNING *; 

