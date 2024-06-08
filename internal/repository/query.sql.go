// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkWorkspaceExists = `-- name: CheckWorkspaceExists :one
SELECT id, name, description, created_at, updated_at, bookmark_count, user_count FROM workspace WHERE id = $1
`

func (q *Queries) CheckWorkspaceExists(ctx context.Context, id int64) (Workspace, error) {
	row := q.db.QueryRow(ctx, checkWorkspaceExists, id)
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

const createBookmark = `-- name: CreateBookmark :one
INSERT INTO bookmark (link, workspace_id, category_id, summary, user_id, title)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, link, category_id, workspace_id, summary, created_at, updated_at, user_id, title, is_read
`

type CreateBookmarkParams struct {
	Link        *string `db:"link" json:"link"`
	WorkspaceID *int64  `db:"workspace_id" json:"workspace_id"`
	CategoryID  *int64  `db:"category_id" json:"category_id"`
	Summary     *string `db:"summary" json:"summary"`
	UserID      *int64  `db:"user_id" json:"user_id"`
	Title       *string `db:"title" json:"title"`
}

func (q *Queries) CreateBookmark(ctx context.Context, arg CreateBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRow(ctx, createBookmark,
		arg.Link,
		arg.WorkspaceID,
		arg.CategoryID,
		arg.Summary,
		arg.UserID,
		arg.Title,
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
		&i.UserID,
		&i.Title,
		&i.IsRead,
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

const createDeviceInfo = `-- name: CreateDeviceInfo :one
INSERT INTO device_info (user_id, registration_token)
VALUES ($1, $2)
RETURNING id, user_id, registration_token, created_at, updated_at
`

type CreateDeviceInfoParams struct {
	UserID            *int64  `db:"user_id" json:"user_id"`
	RegistrationToken *string `db:"registration_token" json:"registration_token"`
}

func (q *Queries) CreateDeviceInfo(ctx context.Context, arg CreateDeviceInfoParams) (DeviceInfo, error) {
	row := q.db.QueryRow(ctx, createDeviceInfo, arg.UserID, arg.RegistrationToken)
	var i DeviceInfo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RegistrationToken,
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

const createNotificationHistory = `-- name: CreateNotificationHistory :one
INSERT INTO notification_history (user_id, notification_title)
VALUES ($1, $2)
RETURNING id, user_id, notification_title, is_read, created_at, updated_at
`

type CreateNotificationHistoryParams struct {
	UserID            *int64  `db:"user_id" json:"user_id"`
	NotificationTitle *string `db:"notification_title" json:"notification_title"`
}

func (q *Queries) CreateNotificationHistory(ctx context.Context, arg CreateNotificationHistoryParams) (NotificationHistory, error) {
	row := q.db.QueryRow(ctx, createNotificationHistory, arg.UserID, arg.NotificationTitle)
	var i NotificationHistory
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.NotificationTitle,
		&i.IsRead,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createRecommendLink = `-- name: CreateRecommendLink :one
INSERT INTO recommend_link (workspace_id, category_id, link, title)
VALUES ($1, $2, $3, $4)
RETURNING id, workspace_id, link, created_at, updated_at, category_id, title
`

type CreateRecommendLinkParams struct {
	WorkspaceID *int64  `db:"workspace_id" json:"workspace_id"`
	CategoryID  *int64  `db:"category_id" json:"category_id"`
	Link        *string `db:"link" json:"link"`
	Title       *string `db:"title" json:"title"`
}

func (q *Queries) CreateRecommendLink(ctx context.Context, arg CreateRecommendLinkParams) (RecommendLink, error) {
	row := q.db.QueryRow(ctx, createRecommendLink,
		arg.WorkspaceID,
		arg.CategoryID,
		arg.Link,
		arg.Title,
	)
	var i RecommendLink
	err := row.Scan(
		&i.ID,
		&i.WorkspaceID,
		&i.Link,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CategoryID,
		&i.Title,
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
VALUES ($1, $2 )
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

const createWorkspaceCode = `-- name: CreateWorkspaceCode :one
INSERT INTO workspace_code (workspace_id, code, user_id)
VALUES ($1, $2, $3)
RETURNING id, workspace_id, code, created_at, updated_at, user_id
`

type CreateWorkspaceCodeParams struct {
	WorkspaceID *int64  `db:"workspace_id" json:"workspace_id"`
	Code        *string `db:"code" json:"code"`
	UserID      *int64  `db:"user_id" json:"user_id"`
}

func (q *Queries) CreateWorkspaceCode(ctx context.Context, arg CreateWorkspaceCodeParams) (WorkspaceCode, error) {
	row := q.db.QueryRow(ctx, createWorkspaceCode, arg.WorkspaceID, arg.Code, arg.UserID)
	var i WorkspaceCode
	err := row.Scan(
		&i.ID,
		&i.WorkspaceID,
		&i.Code,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
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

const exitWorkspace = `-- name: ExitWorkspace :exec
DELETE FROM workspace_user WHERE workspace_id = $1 AND user_id = $2
`

type ExitWorkspaceParams struct {
	WorkspaceID int64 `db:"workspace_id" json:"workspace_id"`
	UserID      int64 `db:"user_id" json:"user_id"`
}

func (q *Queries) ExitWorkspace(ctx context.Context, arg ExitWorkspaceParams) error {
	_, err := q.db.Exec(ctx, exitWorkspace, arg.WorkspaceID, arg.UserID)
	return err
}

const findBookmark = `-- name: FindBookmark :one
SELECT bookmark.id, bookmark.link, bookmark.category_id, bookmark.workspace_id, bookmark.summary, bookmark.created_at, bookmark.updated_at, bookmark.user_id, bookmark.title, bookmark.is_read, workspace.id, workspace.name, workspace.description, workspace.created_at, workspace.updated_at, workspace.bookmark_count, workspace.user_count, category.id, category.name, category.created_at, category.updated_at FROM bookmark
JOIN workspace on workspace.id = bookmark.workspace_id
JOIN category on category.id = bookmark.category_id
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
		&i.Bookmark.UserID,
		&i.Bookmark.Title,
		&i.Bookmark.IsRead,
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

func (q *Queries) FindBookmarkComment(ctx context.Context, id int64) ([]*BookmarkCommentRow, error) {
	row := q.db.QueryRow(ctx, findBookmarkComment, id)
	var comments []*BookmarkCommentRow
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

const findCategoryByName = `-- name: FindCategoryByName :one
SELECT id, name, created_at, updated_at FROM category WHERE name = $1
`

func (q *Queries) FindCategoryByName(ctx context.Context, name *string) (Category, error) {
	row := q.db.QueryRow(ctx, findCategoryByName, name)
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

const findDeviceInfoByToken = `-- name: FindDeviceInfoByToken :one
SELECT id, user_id, registration_token, created_at, updated_at FROM device_info WHERE registration_token = $1
`

func (q *Queries) FindDeviceInfoByToken(ctx context.Context, registrationToken *string) (DeviceInfo, error) {
	row := q.db.QueryRow(ctx, findDeviceInfoByToken, registrationToken)
	var i DeviceInfo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RegistrationToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findDuplicateBookmark = `-- name: FindDuplicateBookmark :one
SELECT id FROM bookmark WHERE workspace_id = $1 AND "link" = $2
`

type FindDuplicateBookmarkParams struct {
	WorkspaceID *int64  `db:"workspace_id" json:"workspace_id"`
	Link        *string `db:"link" json:"link"`
}

func (q *Queries) FindDuplicateBookmark(ctx context.Context, arg FindDuplicateBookmarkParams) (int64, error) {
	row := q.db.QueryRow(ctx, findDuplicateBookmark, arg.WorkspaceID, arg.Link)
	var id int64
	err := row.Scan(&id)
	return id, err
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

const findRecommendLinks = `-- name: FindRecommendLinks :many
SELECT c.name, recommend_links FROM top_workspace_categories JOIN category c ON c.id = top_workspace_categories.category_id
WHERE workspace_id = $1
`

type FindRecommendLinksRow struct {
	Name           *string         `db:"name" json:"name"`
	RecommendLinks []RecommendLink `db:"recommend_links" json:"recommend_links"`
}

func (q *Queries) FindRecommendLinks(ctx context.Context, workspaceID *int64) ([]FindRecommendLinksRow, error) {
	rows, err := q.db.Query(ctx, findRecommendLinks, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindRecommendLinksRow
	for rows.Next() {
		var i FindRecommendLinksRow
		if err := rows.Scan(&i.Name, &i.RecommendLinks); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
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

const findTopCategories = `-- name: FindTopCategories :many
SELECT
	category.id,
  category.name,
  COUNT(*) AS bookmark_count
FROM
    bookmark
JOIN category ON category.id = bookmark.category_id
WHERE workspace_id = $1
GROUP BY
    workspace_id, category_id, category.name, category.id
ORDER BY
    workspace_id, bookmark_count DESC
LIMIT 3
`

type FindTopCategoriesRow struct {
	ID            int64   `db:"id" json:"id"`
	Name          *string `db:"name" json:"name"`
	BookmarkCount int64   `db:"bookmark_count" json:"bookmark_count"`
}

func (q *Queries) FindTopCategories(ctx context.Context, workspaceID *int64) ([]FindTopCategoriesRow, error) {
	rows, err := q.db.Query(ctx, findTopCategories, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindTopCategoriesRow
	for rows.Next() {
		var i FindTopCategoriesRow
		if err := rows.Scan(&i.ID, &i.Name, &i.BookmarkCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findUnreadBookmark = `-- name: FindUnreadBookmark :many
SELECT u.id, u.username, w.name AS workspace_name, bookmarks, device_infos FROM unread_bookmark 
JOIN workspace w ON w.id = unread_bookmark.workspace_id
JOIN "user" u ON u.id = unread_bookmark.user_id
`

type FindUnreadBookmarkRow struct {
	ID            int64        `db:"id" json:"id"`
	Username      *string      `db:"username" json:"username"`
	WorkspaceName *string      `db:"workspace_name" json:"workspace_name"`
	Bookmarks     []Bookmark   `db:"bookmarks" json:"bookmarks"`
	DeviceInfos   []DeviceInfo `db:"device_infos" json:"device_infos"`
}

func (q *Queries) FindUnreadBookmark(ctx context.Context) ([]FindUnreadBookmarkRow, error) {
	rows, err := q.db.Query(ctx, findUnreadBookmark)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindUnreadBookmarkRow
	for rows.Next() {
		var i FindUnreadBookmarkRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.WorkspaceName,
			&i.Bookmarks,
			&i.DeviceInfos,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findUnreadNotificationHistory = `-- name: FindUnreadNotificationHistory :one
SELECT id, username, notifications FROM unread_notifications WHERE id = $1
`

func (q *Queries) FindUnreadNotificationHistory(ctx context.Context, id int64) (UnreadNotification, error) {
	row := q.db.QueryRow(ctx, findUnreadNotificationHistory, id)
	var i UnreadNotification
	err := row.Scan(&i.ID, &i.Username, &i.Notifications)
	return i, err
}

const findUserById = `-- name: FindUserById :one
SELECT "user".id, "user".username, "user".provider, "user".created_at, "user".updated_at  FROM "user" WHERE "id" = $1 LIMIT 1
`

type FindUserByIdRow struct {
	ID        int64              `db:"id" json:"id"`
	Username  *string            `db:"username" json:"username"`
	Provider  *string            `db:"provider" json:"provider"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

func (q *Queries) FindUserById(ctx context.Context, id int64) (FindUserByIdRow, error) {
	row := q.db.QueryRow(ctx, findUserById, id)
	var i FindUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Provider,
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
SELECT id,categories,users FROM workspace_user_category WHERE id = $1
`

type FindWorkspaceRow struct {
	ID         int64              `db:"id" json:"id"`
	Categories []*Category        `db:"categories" json:"categories"`
	Users      []*FindUserByIdRow `db:"users" json:"users"`
}

func (q *Queries) FindWorkspace(ctx context.Context, id int64) (FindWorkspaceRow, error) {
	row := q.db.QueryRow(ctx, findWorkspace, id)
	var i FindWorkspaceRow
	err := row.Scan(&i.ID, &i.Categories, &i.Users)
	return i, err
}

const findWorkspaceCategory = `-- name: FindWorkspaceCategory :one
SELECT categories FROM workspace_category_list WHERE id = $1
`

func (q *Queries) FindWorkspaceCategory(ctx context.Context, id int64) ([]*Category, error) {
	row := q.db.QueryRow(ctx, findWorkspaceCategory, id)
	var categories []*Category
	err := row.Scan(&categories)
	return categories, err
}

const findWorkspaceCategoryBookmark = `-- name: FindWorkspaceCategoryBookmark :many
SELECT id, link, category_id, workspace_id, summary, created_at, updated_at, user_id, title, is_read FROM bookmark WHERE workspace_id = $1 AND category_id = $2
`

type FindWorkspaceCategoryBookmarkParams struct {
	WorkspaceID *int64 `db:"workspace_id" json:"workspace_id"`
	CategoryID  *int64 `db:"category_id" json:"category_id"`
}

func (q *Queries) FindWorkspaceCategoryBookmark(ctx context.Context, arg FindWorkspaceCategoryBookmarkParams) ([]Bookmark, error) {
	rows, err := q.db.Query(ctx, findWorkspaceCategoryBookmark, arg.WorkspaceID, arg.CategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bookmark
	for rows.Next() {
		var i Bookmark
		if err := rows.Scan(
			&i.ID,
			&i.Link,
			&i.CategoryID,
			&i.WorkspaceID,
			&i.Summary,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.Title,
			&i.IsRead,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findWorkspaceCode = `-- name: FindWorkspaceCode :one
SELECT workspace_code.id, workspace_code.workspace_id, workspace_code.code, workspace_code.created_at, workspace_code.updated_at, workspace_code.user_id, workspace.id, workspace.name, workspace.description, workspace.created_at, workspace.updated_at, workspace.bookmark_count, workspace.user_count FROM workspace_code
JOIN workspace ON workspace_code.workspace_id = workspace.id
WHERE workspace_code.code = $1
`

type FindWorkspaceCodeRow struct {
	WorkspaceCode WorkspaceCode `db:"workspace_code" json:"workspace_code"`
	Workspace     Workspace     `db:"workspace" json:"workspace"`
}

func (q *Queries) FindWorkspaceCode(ctx context.Context, code *string) (FindWorkspaceCodeRow, error) {
	row := q.db.QueryRow(ctx, findWorkspaceCode, code)
	var i FindWorkspaceCodeRow
	err := row.Scan(
		&i.WorkspaceCode.ID,
		&i.WorkspaceCode.WorkspaceID,
		&i.WorkspaceCode.Code,
		&i.WorkspaceCode.CreatedAt,
		&i.WorkspaceCode.UpdatedAt,
		&i.WorkspaceCode.UserID,
		&i.Workspace.ID,
		&i.Workspace.Name,
		&i.Workspace.Description,
		&i.Workspace.CreatedAt,
		&i.Workspace.UpdatedAt,
		&i.Workspace.BookmarkCount,
		&i.Workspace.UserCount,
	)
	return i, err
}

const findWorkspaceCodeByWorkspaceID = `-- name: FindWorkspaceCodeByWorkspaceID :one
SELECT id, workspace_id, code, created_at, updated_at, user_id FROM workspace_code WHERE workspace_id = $1
`

func (q *Queries) FindWorkspaceCodeByWorkspaceID(ctx context.Context, workspaceID *int64) (WorkspaceCode, error) {
	row := q.db.QueryRow(ctx, findWorkspaceCodeByWorkspaceID, workspaceID)
	var i WorkspaceCode
	err := row.Scan(
		&i.ID,
		&i.WorkspaceID,
		&i.Code,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const findWorkspaceCodeByWorkspaceIDAndUserID = `-- name: FindWorkspaceCodeByWorkspaceIDAndUserID :one
SELECT id, workspace_id, code, created_at, updated_at, user_id FROM workspace_code WHERE workspace_id = $1 AND user_id = $2
`

type FindWorkspaceCodeByWorkspaceIDAndUserIDParams struct {
	WorkspaceID *int64 `db:"workspace_id" json:"workspace_id"`
	UserID      *int64 `db:"user_id" json:"user_id"`
}

func (q *Queries) FindWorkspaceCodeByWorkspaceIDAndUserID(ctx context.Context, arg FindWorkspaceCodeByWorkspaceIDAndUserIDParams) (WorkspaceCode, error) {
	row := q.db.QueryRow(ctx, findWorkspaceCodeByWorkspaceIDAndUserID, arg.WorkspaceID, arg.UserID)
	var i WorkspaceCode
	err := row.Scan(
		&i.ID,
		&i.WorkspaceID,
		&i.Code,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const findWorkspaceInfo = `-- name: FindWorkspaceInfo :one
SELECT w.name, w.description, user_bookmark_count 
FROM workspace_user_bookmark_count wu 
JOIN workspace w ON wu.workspace_id = w.id
WHERE workspace_id = $1
`

type FindWorkspaceInfoRow struct {
	Name              *string             `db:"name" json:"name"`
	Description       *string             `db:"description" json:"description"`
	UserBookmarkCount []UserBookmarkCount `db:"user_bookmark_count" json:"user_bookmark_count"`
}

func (q *Queries) FindWorkspaceInfo(ctx context.Context, workspaceID int64) (FindWorkspaceInfoRow, error) {
	row := q.db.QueryRow(ctx, findWorkspaceInfo, workspaceID)
	var i FindWorkspaceInfoRow
	err := row.Scan(&i.Name, &i.Description, &i.UserBookmarkCount)
	return i, err
}

const findWorkspaceJoinedUser = `-- name: FindWorkspaceJoinedUser :one
SELECT workspace_id, user_id FROM workspace_user WHERE workspace_id = $1 AND user_id = $2
`

type FindWorkspaceJoinedUserParams struct {
	WorkspaceID int64 `db:"workspace_id" json:"workspace_id"`
	UserID      int64 `db:"user_id" json:"user_id"`
}

func (q *Queries) FindWorkspaceJoinedUser(ctx context.Context, arg FindWorkspaceJoinedUserParams) (WorkspaceUser, error) {
	row := q.db.QueryRow(ctx, findWorkspaceJoinedUser, arg.WorkspaceID, arg.UserID)
	var i WorkspaceUser
	err := row.Scan(&i.WorkspaceID, &i.UserID)
	return i, err
}

const isUserJoinedWorkspace = `-- name: IsUserJoinedWorkspace :one
SELECT workspace_id, user_id FROM workspace_user WHERE workspace_id = $1 AND user_id = $2
`

type IsUserJoinedWorkspaceParams struct {
	WorkspaceID int64 `db:"workspace_id" json:"workspace_id"`
	UserID      int64 `db:"user_id" json:"user_id"`
}

func (q *Queries) IsUserJoinedWorkspace(ctx context.Context, arg IsUserJoinedWorkspaceParams) (WorkspaceUser, error) {
	row := q.db.QueryRow(ctx, isUserJoinedWorkspace, arg.WorkspaceID, arg.UserID)
	var i WorkspaceUser
	err := row.Scan(&i.WorkspaceID, &i.UserID)
	return i, err
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

const readBookmark = `-- name: ReadBookmark :exec
UPDATE bookmark
SET
  is_read = true,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) ReadBookmark(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, readBookmark, id)
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

const searchWorkspaceBookmark = `-- name: SearchWorkspaceBookmark :many
SELECT bookmark.id, bookmark.link, bookmark.category_id, bookmark.workspace_id, bookmark.summary, bookmark.created_at, bookmark.updated_at, bookmark.user_id, bookmark.title, bookmark.is_read, category.name AS category_name
FROM bookmark
JOIN category ON bookmark.category_id = category.id
WHERE workspace_id = $1
  AND (array_length($2::bigint[], 1) IS NULL OR array_length($2::bigint[], 1) = 0 OR user_id = ANY($2::bigint[]))
  AND (array_length($3::bigint[], 1) IS NULL OR array_length($3::bigint[], 1) = 0 OR category_id = ANY($3::bigint[]))
`

type SearchWorkspaceBookmarkParams struct {
	WorkspaceID *int64  `db:"workspace_id" json:"workspace_id"`
	UserIds     []int64 `db:"user_ids" json:"user_ids"`
	CategoryIds []int64 `db:"category_ids" json:"category_ids"`
}

type SearchWorkspaceBookmarkRow struct {
	ID           int64              `db:"id" json:"id"`
	Link         *string            `db:"link" json:"link"`
	CategoryID   *int64             `db:"category_id" json:"category_id"`
	WorkspaceID  *int64             `db:"workspace_id" json:"workspace_id"`
	Summary      *string            `db:"summary" json:"summary"`
	CreatedAt    pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	UserID       *int64             `db:"user_id" json:"user_id"`
	Title        *string            `db:"title" json:"title"`
	IsRead       *bool              `db:"is_read" json:"is_read"`
	CategoryName *string            `db:"category_name" json:"category_name"`
}

func (q *Queries) SearchWorkspaceBookmark(ctx context.Context, arg SearchWorkspaceBookmarkParams) ([]SearchWorkspaceBookmarkRow, error) {
	rows, err := q.db.Query(ctx, searchWorkspaceBookmark, arg.WorkspaceID, arg.UserIds, arg.CategoryIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchWorkspaceBookmarkRow
	for rows.Next() {
		var i SearchWorkspaceBookmarkRow
		if err := rows.Scan(
			&i.ID,
			&i.Link,
			&i.CategoryID,
			&i.WorkspaceID,
			&i.Summary,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.Title,
			&i.IsRead,
			&i.CategoryName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBookmark = `-- name: UpdateBookmark :one
UPDATE bookmark 
SET
  link = coalesce($2,link),
  workspace_id = coalesce($3,workspace_id),
  category_id = coalesce($4,category_id),
  summary = coalesce($5,summary),
  title = coalesce($6,summary),
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, link, category_id, workspace_id, summary, created_at, updated_at, user_id, title, is_read
`

type UpdateBookmarkParams struct {
	ID          int64   `db:"id" json:"id"`
	Link        *string `db:"link" json:"link"`
	WorkspaceID *int64  `db:"workspace_id" json:"workspace_id"`
	CategoryID  *int64  `db:"category_id" json:"category_id"`
	Summary     *string `db:"summary" json:"summary"`
	Title       *string `db:"title" json:"title"`
}

func (q *Queries) UpdateBookmark(ctx context.Context, arg UpdateBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRow(ctx, updateBookmark,
		arg.ID,
		arg.Link,
		arg.WorkspaceID,
		arg.CategoryID,
		arg.Summary,
		arg.Title,
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
		&i.UserID,
		&i.Title,
		&i.IsRead,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE category
SET
  name = coalesce($1,name),
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
  comment_context = coalesce($1,comment_context),
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

const updateInviteCode = `-- name: UpdateInviteCode :one
UPDATE "invite_code"
SET
  code = coalesce($1,code),
  updated_at = CURRENT_TIMESTAMP
WHERE id = $2
RETURNING id, workspace_id, code, expired_at, created_at, updated_at
`

type UpdateInviteCodeParams struct {
	Code *string `db:"code" json:"code"`
	ID   int64   `db:"id" json:"id"`
}

func (q *Queries) UpdateInviteCode(ctx context.Context, arg UpdateInviteCodeParams) (InviteCode, error) {
	row := q.db.QueryRow(ctx, updateInviteCode, arg.Code, arg.ID)
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

const updateRefreshToken = `-- name: UpdateRefreshToken :one
UPDATE refresh_token
SET
    token = coalesce($1,token),
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
    refresh_token = coalesce($1,refresh_token),
    username = coalesce($2,username),
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
  name = coalesce($1,name),
  description = coalesce($2,description),
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

const updateWorkspaceCode = `-- name: UpdateWorkspaceCode :one
UPDATE workspace_code
SET
  code = coalesce($1,code),
  updated_at = CURRENT_TIMESTAMP
WHERE id = $2
RETURNING id, workspace_id, code, created_at, updated_at, user_id
`

type UpdateWorkspaceCodeParams struct {
	Code *string `db:"code" json:"code"`
	ID   int64   `db:"id" json:"id"`
}

func (q *Queries) UpdateWorkspaceCode(ctx context.Context, arg UpdateWorkspaceCodeParams) (WorkspaceCode, error) {
	row := q.db.QueryRow(ctx, updateWorkspaceCode, arg.Code, arg.ID)
	var i WorkspaceCode
	err := row.Scan(
		&i.ID,
		&i.WorkspaceID,
		&i.Code,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}
