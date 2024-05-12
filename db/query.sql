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

-- name: CreateWorkspace :one
INSERT INTO workspace (name)
VALUES ($1)
RETURNING *;

-- name: UpdateWorkspace :one
UPDATE workspace
SET
  name = $1,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $2
RETURNING *;

-- name: DeleteWorkspace :exec
DELETE FROM workspace WHERE id = $1;

-- name: FindWorkspace :one
SELECT * FROM workspace_user_category WHERE id = $1;

-- name: FindInviteCodeByWorkspaceID :one
SELECT * FROM invite_code WHERE workspace_id = $1;

-- name: CreateInviteCode :one
INSERT INTO invite_code (workspace_id, code)
VALUES ($1, $2)
RETURNING *;


