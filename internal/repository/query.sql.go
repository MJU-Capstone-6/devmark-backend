// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package repository

import (
	"context"
)

const createBookmark = `-- name: CreateBookmark :one
INSERT INTO bookmark (link, workspace_id, category_id, summary)
VALUES ($1, $2, $3, $4)
RETURNING id, link, category_id, workspace_id, summary, created_at, updated_at
`

type CreateBookmarkParams struct {
	Link        *string `db:"link" json:"link"`
	WorkspaceID *int64  `db:"workspace_id" json:"workspace_id"`
	CategoryID  *int64  `db:"category_id" json:"category_id"`
	Summary     *string `db:"summary" json:"summary"`
}

func (q *Queries) CreateBookmark(ctx context.Context, arg CreateBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRow(ctx, createBookmark,
		arg.Link,
		arg.WorkspaceID,
		arg.CategoryID,
		arg.Summary,
	)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.Link,
		&i.CategoryID,
		&i.WorkspaceID,
		&i.Summary,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createCategory = `-- name: CreateCategory :one
INSERT INTO category (name)
VALUES ($1)
RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateCategory(ctx context.Context, name *string) (Category, error) {
	row := q.db.QueryRow(ctx, createCategory, name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createComment = `-- name: CreateComment :one
INSERT INTO "comment" (bookmark_id, user_id, comment_context)
VALUES ($1, $2, $3)
RETURNING id, bookmark_id, user_id, comment_context, created_at, updated_at
`

type CreateCommentParams struct {
	BookmarkID     *int64  `db:"bookmark_id" json:"bookmark_id"`
	UserID         *int64  `db:"user_id" json:"user_id"`
	CommentContext *string `db:"comment_context" json:"comment_context"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, createComment, arg.BookmarkID, arg.UserID, arg.CommentContext)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.BookmarkID,
		&i.UserID,
		&i.CommentContext,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createInviteCode = `-- name: CreateInviteCode :one
INSERT INTO invite_code (workspace_id, code)
VALUES ($1, $2)
RETURNING id, workspace_id, code, expired_at, created_at, updated_at
`

type CreateInviteCodeParams struct {
	WorkspaceID *int32  `db:"workspace_id" json:"workspace_id"`
	Code        *string `db:"code" json:"code"`
}

func (q *Queries) CreateInviteCode(ctx context.Context, arg CreateInviteCodeParams) (InviteCode, error) {
	row := q.db.QueryRow(ctx, createInviteCode, arg.WorkspaceID, arg.Code)
	var i InviteCode
	err := row.Scan(
		&i.ID,
		&i.WorkspaceID,
		&i.Code,
		&i.ExpiredAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createRefreshToken = `-- name: CreateRefreshToken :one
INSERT INTO refresh_token (token, user_id)
VALUES ($1, $2)
RETURNING id, token, user_id, created_at, updated_at
`

type CreateRefreshTokenParams struct {
	Token  *string `db:"token" json:"token"`
	UserID *int32  `db:"user_id" json:"user_id"`
}

func (q *Queries) CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRow(ctx, createRefreshToken, arg.Token, arg.UserID)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (username, provider)
VALUES ($1, $2)
RETURNING id, username, provider, refresh_token, created_at, updated_at
`

type CreateUserParams struct {
	Username *string `db:"username" json:"username"`
	Provider *string `db:"provider" json:"provider"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Provider)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Provider,
		&i.RefreshToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createWorkspace = `-- name: CreateWorkspace :one
INSERT INTO workspace (name, description)
VALUES ($1, $2)
RETURNING id, name, description, created_at, updated_at, bookmark_count, user_count
`

type CreateWorkspaceParams struct {
	Name        *string `db:"name" json:"name"`
	Description *string `db:"description" json:"description"`
}

func (q *Queries) CreateWorkspace(ctx context.Context, arg CreateWorkspaceParams) (Workspace, error) {
	row := q.db.QueryRow(ctx, createWorkspace, arg.Name, arg.Description)
	var i Workspace
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BookmarkCount,
		&i.UserCount,
	)
	return i, err
}

const deleteBookmark = `-- name: DeleteBookmark :exec
DELETE FROM bookmark WHERE id = $1
`

func (q *Queries) DeleteBookmark(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteBookmark, id)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCategory, id)
	return err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM "comment" WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteComment, id)
	return err
}

const deleteWorkspace = `-- name: DeleteWorkspace :exec
DELETE FROM workspace WHERE id = $1
`

func (q *Queries) DeleteWorkspace(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteWorkspace, id)
	return err
}

const findBookmark = `-- name: FindBookmark :one
SELECT bookmark.id, bookmark.link, bookmark.category_id, bookmark.workspace_id, bookmark.summary, bookmark.created_at, bookmark.updated_at, workspace.id, workspace.name, workspace.description, workspace.created_at, workspace.updated_at, workspace.bookmark_count, workspace.user_count, category.id, category.name, category.created_at, category.updated_at FROM bookmark
JOIN workspace on workspace.id = bookmark.workspace_id
JOIN category on category.id = bookmark.workspace_id
WHERE bookmark.id = $1
`

type FindBookmarkRow struct {
	Bookmark  Bookmark  `db:"bookmark" json:"bookmark"`
	Workspace Workspace `db:"workspace" json:"workspace"`
	Category  Category  `db:"category" json:"category"`
}

func (q *Queries) FindBookmark(ctx context.Context, id int64) (FindBookmarkRow, error) {
	row := q.db.QueryRow(ctx, findBookmark, id)
	var i FindBookmarkRow
	err := row.Scan(
		&i.Bookmark.ID,
		&i.Bookmark.Link,
		&i.Bookmark.CategoryID,
		&i.Bookmark.WorkspaceID,
		&i.Bookmark.Summary,
		&i.Bookmark.CreatedAt,
		&i.Bookmark.UpdatedAt,
		&i.Workspace.ID,
		&i.Workspace.Name,
		&i.Workspace.Description,
		&i.Workspace.CreatedAt,
		&i.Workspace.UpdatedAt,
		&i.Workspace.BookmarkCount,
		&i.Workspace.UserCount,
		&i.Category.ID,
		&i.Category.Name,
		&i.Category.CreatedAt,
		&i.Category.UpdatedAt,
	)
	return i, err
}

const findBookmarkComment = `-- name: FindBookmarkComment :one
SELECT comments FROM bookmark_comment WHERE id = $1
`

func (q *Queries) FindBookmarkComment(ctx context.Context, id int64) ([]Comment, error) {
	row := q.db.QueryRow(ctx, findBookmarkComment, id)
	var comments []Comment
	err := row.Scan(&comments)
	return comments, err
}

const findCategoryById = `-- name: FindCategoryById :one
SELECT id, name, created_at, updated_at FROM category WHERE id = $1
`

func (q *Queries) FindCategoryById(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRow(ctx, findCategoryById, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findComment = `-- name: FindComment :one
SELECT id, bookmark_id, user_id, comment_context, created_at, updated_at FROM "comment" WHERE id = $1
`

func (q *Queries) FindComment(ctx context.Context, id int64) (Comment, error) {
	row := q.db.QueryRow(ctx, findComment, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.BookmarkID,
		&i.UserID,
		&i.CommentContext,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findInviteCodeByCode = `-- name: FindInviteCodeByCode :one
SELECT id, workspace_id, code, expired_at, created_at, updated_at FROM invite_code WHERE code = $1
`

func (q *Queries) FindInviteCodeByCode(ctx context.Context, code *string) (InviteCode, error) {
	row := q.db.QueryRow(ctx, findInviteCodeByCode, code)
	var i InviteCode
	err := row.Scan(
		&i.ID,
		&i.WorkspaceID,
		&i.Code,
		&i.ExpiredAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findInviteCodeByWorkspaceID = `-- name: FindInviteCodeByWorkspaceID :one
SELECT id, workspace_id, code, expired_at, created_at, updated_at FROM invite_code WHERE workspace_id = $1
`

func (q *Queries) FindInviteCodeByWorkspaceID(ctx context.Context, workspaceID *int32) (InviteCode, error) {
	row := q.db.QueryRow(ctx, findInviteCodeByWorkspaceID, workspaceID)
	var i InviteCode
	err := row.Scan(
		&i.ID,
		&i.WorkspaceID,
		&i.Code,
		&i.ExpiredAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findRefreshTokenByUserID = `-- name: FindRefreshTokenByUserID :one
SELECT id, token, user_id, created_at, updated_at from "refresh_token" WHERE "user_id" = $1 LIMIT 1
`

func (q *Queries) FindRefreshTokenByUserID(ctx context.Context, userID *int32) (RefreshToken, error) {
	row := q.db.QueryRow(ctx, findRefreshTokenByUserID, userID)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserById = `-- name: FindUserById :one
SELECT id, username, provider, refresh_token, created_at, updated_at FROM "user" WHERE "id" = $1 LIMIT 1
`

func (q *Queries) FindUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, findUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Provider,
		&i.RefreshToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByUsername = `-- name: FindUserByUsername :one
SELECT id, username, provider, refresh_token, created_at, updated_at FROM "user" WHERE "username" = $1 LIMIT 1
`

func (q *Queries) FindUserByUsername(ctx context.Context, username *string) (User, error) {
	row := q.db.QueryRow(ctx, findUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Provider,
		&i.RefreshToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserWorkspace = `-- name: FindUserWorkspace :one
SELECT id, workspaces FROM user_workspace_view WHERE id = $1
`

func (q *Queries) FindUserWorkspace(ctx context.Context, id *int64) (UserWorkspaceView, error) {
	row := q.db.QueryRow(ctx, findUserWorkspace, id)
	var i UserWorkspaceView
	err := row.Scan(&i.ID, &i.Workspaces)
	return i, err
}

const findWorkspace = `-- name: FindWorkspace :one
SELECT id, name, description, created_at, updated_at, bookmark_count, user_count, categories, users FROM workspace_user_category WHERE id = $1
`

func (q *Queries) FindWorkspace(ctx context.Context, id int64) (WorkspaceUserCategory, error) {
	row := q.db.QueryRow(ctx, findWorkspace, id)
	var i WorkspaceUserCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BookmarkCount,
		&i.UserCount,
		&i.Categories,
		&i.Users,
	)
	return i, err
}

const findWorkspaceCategory = `-- name: FindWorkspaceCategory :one
SELECT categories FROM workspace_category_list WHERE id = $1
`

func (q *Queries) FindWorkspaceCategory(ctx context.Context, id int64) ([]Category, error) {
	row := q.db.QueryRow(ctx, findWorkspaceCategory, id)
	var categories []Category
	err := row.Scan(&categories)
	return categories, err
}

const joinWorkspace = `-- name: JoinWorkspace :exec
INSERT INTO workspace_user (workspace_id, user_id)
VALUES ($1, $2)
RETURNING workspace_id, user_id
`

type JoinWorkspaceParams struct {
	WorkspaceID int64 `db:"workspace_id" json:"workspace_id"`
	UserID      int64 `db:"user_id" json:"user_id"`
}

func (q *Queries) JoinWorkspace(ctx context.Context, arg JoinWorkspaceParams) error {
	_, err := q.db.Exec(ctx, joinWorkspace, arg.WorkspaceID, arg.UserID)
	return err
}

const joinWorkspaceWithoutCode = `-- name: JoinWorkspaceWithoutCode :exec
INSERT INTO workspace_user (workspace_id, user_id)
VALUES ($1, $2)
`

type JoinWorkspaceWithoutCodeParams struct {
	WorkspaceID int64 `db:"workspace_id" json:"workspace_id"`
	UserID      int64 `db:"user_id" json:"user_id"`
}

func (q *Queries) JoinWorkspaceWithoutCode(ctx context.Context, arg JoinWorkspaceWithoutCodeParams) error {
	_, err := q.db.Exec(ctx, joinWorkspaceWithoutCode, arg.WorkspaceID, arg.UserID)
	return err
}

const registerCategoryToWorkspace = `-- name: RegisterCategoryToWorkspace :exec
INSERT INTO workspace_category (workspace_id, category_id)
VALUES ($1, $2)
RETURNING workspace_id, category_id
`

type RegisterCategoryToWorkspaceParams struct {
	WorkspaceID int64 `db:"workspace_id" json:"workspace_id"`
	CategoryID  int64 `db:"category_id" json:"category_id"`
}

func (q *Queries) RegisterCategoryToWorkspace(ctx context.Context, arg RegisterCategoryToWorkspaceParams) error {
	_, err := q.db.Exec(ctx, registerCategoryToWorkspace, arg.WorkspaceID, arg.CategoryID)
	return err
}

const updateBookmark = `-- name: UpdateBookmark :one
UPDATE bookmark 
SET
  link = $2,
  workspace_id = $3,
  category_id = $4,
  summary = $5,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, link, category_id, workspace_id, summary, created_at, updated_at
`

type UpdateBookmarkParams struct {
	ID          int64   `db:"id" json:"id"`
	Link        *string `db:"link" json:"link"`
	WorkspaceID *int64  `db:"workspace_id" json:"workspace_id"`
	CategoryID  *int64  `db:"category_id" json:"category_id"`
	Summary     *string `db:"summary" json:"summary"`
}

func (q *Queries) UpdateBookmark(ctx context.Context, arg UpdateBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRow(ctx, updateBookmark,
		arg.ID,
		arg.Link,
		arg.WorkspaceID,
		arg.CategoryID,
		arg.Summary,
	)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.Link,
		&i.CategoryID,
		&i.WorkspaceID,
		&i.Summary,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE category
SET
  name = $1,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $2
RETURNING id, name, created_at, updated_at
`

type UpdateCategoryParams struct {
	Name *string `db:"name" json:"name"`
	ID   int64   `db:"id" json:"id"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, updateCategory, arg.Name, arg.ID)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateComment = `-- name: UpdateComment :one
UPDATE "comment"
SET
  comment_context = $1,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $2
RETURNING id, bookmark_id, user_id, comment_context, created_at, updated_at
`

type UpdateCommentParams struct {
	CommentContext *string `db:"comment_context" json:"comment_context"`
	ID             int64   `db:"id" json:"id"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, updateComment, arg.CommentContext, arg.ID)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.BookmarkID,
		&i.UserID,
		&i.CommentContext,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateRefreshToken = `-- name: UpdateRefreshToken :one
UPDATE refresh_token
SET
    token = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $2 
RETURNING id, token, user_id, created_at, updated_at
`

type UpdateRefreshTokenParams struct {
	Token *string `db:"token" json:"token"`
	ID    int64   `db:"id" json:"id"`
}

func (q *Queries) UpdateRefreshToken(ctx context.Context, arg UpdateRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRow(ctx, updateRefreshToken, arg.Token, arg.ID)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
Update "user"
SET 
    refresh_token = $1,
    username = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $3
RETURNING id, username, provider, refresh_token, created_at, updated_at
`

type UpdateUserParams struct {
	RefreshToken *int32  `db:"refresh_token" json:"refresh_token"`
	Username     *string `db:"username" json:"username"`
	ID           int64   `db:"id" json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser, arg.RefreshToken, arg.Username, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Provider,
		&i.RefreshToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateWorkspace = `-- name: UpdateWorkspace :one
UPDATE workspace
SET
  name = $1,
  description = $2,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $3
RETURNING id, name, description, created_at, updated_at, bookmark_count, user_count
`

type UpdateWorkspaceParams struct {
	Name        *string `db:"name" json:"name"`
	Description *string `db:"description" json:"description"`
	ID          int64   `db:"id" json:"id"`
}

func (q *Queries) UpdateWorkspace(ctx context.Context, arg UpdateWorkspaceParams) (Workspace, error) {
	row := q.db.QueryRow(ctx, updateWorkspace, arg.Name, arg.Description, arg.ID)
	var i Workspace
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BookmarkCount,
		&i.UserCount,
	)
	return i, err
}
