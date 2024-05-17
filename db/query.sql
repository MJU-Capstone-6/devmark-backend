-- name: FindUserByUsername :one
SELECT * FROM "user" WHERE "username" = $1 LIMIT 1;

-- name: FindUserById :one
SELECT "user".id, "user".username, "user".provider, "user".created_at, "user".updated_at  FROM "user" WHERE "id" = $1 LIMIT 1;


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
INSERT INTO workspace (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateWorkspace :one
UPDATE workspace
SET
  name = $1,
  description = $2,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $3
RETURNING *;

-- name: DeleteWorkspace :exec
DELETE FROM workspace WHERE id = $1;

-- name: FindWorkspace :one
SELECT id,categories,users FROM workspace_user_category WHERE id = $1;

-- name: FindInviteCodeByCode :one
SELECT * FROM invite_code WHERE code = $1;

-- name: FindInviteCodeByWorkspaceID :one
SELECT * FROM invite_code WHERE workspace_id = $1;

-- name: CreateInviteCode :one
INSERT INTO invite_code (workspace_id, code)
VALUES ($1, $2)
RETURNING *;

-- name: JoinWorkspace :exec
INSERT INTO workspace_user (workspace_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: FindCategoryById :one
SELECT * FROM category WHERE id = $1;

-- name: CreateCategory :one
INSERT INTO category (name)
VALUES ($1)
RETURNING *;

-- name: UpdateCategory :one
UPDATE category
SET
  name = $1,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $2
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1;

-- name: RegisterCategoryToWorkspace :exec
INSERT INTO workspace_category (workspace_id, category_id)
VALUES ($1, $2)
RETURNING *;

-- name: FindBookmark :one
SELECT sqlc.embed(bookmark), sqlc.embed(workspace), sqlc.embed(category) FROM bookmark
JOIN workspace on workspace.id = bookmark.workspace_id
JOIN category on category.id = bookmark.workspace_id
WHERE bookmark.id = $1;

-- name: CreateBookmark :one
INSERT INTO bookmark (link, workspace_id, category_id, summary)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateBookmark :one
UPDATE bookmark 
SET
  link = $2,
  workspace_id = $3,
  category_id = $4,
  summary = $5,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteBookmark :exec
DELETE FROM bookmark WHERE id = $1;

-- name: FindUserWorkspace :one
SELECT * FROM user_workspace_view WHERE id = $1;

-- name: JoinWorkspaceWithoutCode :exec
INSERT INTO workspace_user (workspace_id, user_id)
VALUES ($1, $2);

-- name: FindComment :one
SELECT * FROM "comment" WHERE id = $1;

-- name: CreateComment :one
INSERT INTO "comment" (bookmark_id, user_id, comment_context)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateComment :one
UPDATE "comment"
SET
  comment_context = $1,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $2
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM "comment" WHERE id = $1;

-- name: FindBookmarkComment :one
SELECT comments FROM bookmark_comment WHERE id = $1;

-- name: FindWorkspaceCategory :one
SELECT categories FROM workspace_category_list WHERE id = $1;

-- name: FindWorkspaceCategoryBookmark :many
SELECT * FROM bookmark WHERE workspace_id = $1 AND category_id = $2;


