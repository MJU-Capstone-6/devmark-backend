// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bookmark struct {
	ID          int64              `db:"id" json:"id"`
	Link        *string            `db:"link" json:"link"`
	CategoryID  *int64             `db:"category_id" json:"category_id"`
	WorkspaceID *int64             `db:"workspace_id" json:"workspace_id"`
	Summary     *string            `db:"summary" json:"summary"`
	CreatedAt   pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	UserID      *int64             `db:"user_id" json:"user_id"`
	Title       *string            `db:"title" json:"title"`
}

type BookmarkComment struct {
	ID          int64                 `db:"id" json:"id"`
	Link        *string               `db:"link" json:"link"`
	CategoryID  *int64                `db:"category_id" json:"category_id"`
	WorkspaceID *int64                `db:"workspace_id" json:"workspace_id"`
	Summary     *string               `db:"summary" json:"summary"`
	CreatedAt   pgtype.Timestamptz    `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamptz    `db:"updated_at" json:"updated_at"`
	UserID      *int64                `db:"user_id" json:"user_id"`
	Title       *string               `db:"title" json:"title"`
	Comments    []*BookmarkCommentRow `db:"comments" json:"comments"`
}

type Category struct {
	ID        int64              `db:"id" json:"id"`
	Name      *string            `db:"name" json:"name"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

type Comment struct {
	ID             int64              `db:"id" json:"id"`
	BookmarkID     *int64             `db:"bookmark_id" json:"bookmark_id"`
	UserID         *int64             `db:"user_id" json:"user_id"`
	CommentContext *string            `db:"comment_context" json:"comment_context"`
	CreatedAt      pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

type InviteCode struct {
	ID          int64              `db:"id" json:"id"`
	WorkspaceID *int32             `db:"workspace_id" json:"workspace_id"`
	Code        *string            `db:"code" json:"code"`
	ExpiredAt   pgtype.Timestamptz `db:"expired_at" json:"expired_at"`
	CreatedAt   pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

type RefreshToken struct {
	ID        int64              `db:"id" json:"id"`
	Token     *string            `db:"token" json:"token"`
	UserID    *int32             `db:"user_id" json:"user_id"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

type User struct {
	ID           int64              `db:"id" json:"id"`
	Username     *string            `db:"username" json:"username"`
	Provider     *string            `db:"provider" json:"provider"`
	RefreshToken *int32             `db:"refresh_token" json:"refresh_token"`
	CreatedAt    pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

type UserWorkspaceView struct {
	ID         *int64       `db:"id" json:"id"`
	Workspaces []*Workspace `db:"workspaces" json:"workspaces"`
}

type Workspace struct {
	ID            int64              `db:"id" json:"id"`
	Name          *string            `db:"name" json:"name"`
	Description   *string            `db:"description" json:"description"`
	CreatedAt     pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	BookmarkCount *int32             `db:"bookmark_count" json:"bookmark_count"`
	UserCount     *int32             `db:"user_count" json:"user_count"`
}

type WorkspaceCategory struct {
	WorkspaceID int64 `db:"workspace_id" json:"workspace_id"`
	CategoryID  int64 `db:"category_id" json:"category_id"`
}

type WorkspaceCategoryList struct {
	ID            int64              `db:"id" json:"id"`
	Name          *string            `db:"name" json:"name"`
	Description   *string            `db:"description" json:"description"`
	CreatedAt     pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	BookmarkCount *int32             `db:"bookmark_count" json:"bookmark_count"`
	UserCount     *int32             `db:"user_count" json:"user_count"`
	Categories    []*Category        `db:"categories" json:"categories"`
}

type WorkspaceCode struct {
	ID          int64              `db:"id" json:"id"`
	WorkspaceID *int64             `db:"workspace_id" json:"workspace_id"`
	Code        *string            `db:"code" json:"code"`
	CreatedAt   pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

type WorkspaceUser struct {
	WorkspaceID int64 `db:"workspace_id" json:"workspace_id"`
	UserID      int64 `db:"user_id" json:"user_id"`
}

type WorkspaceUserCategory struct {
	ID            int64              `db:"id" json:"id"`
	Name          *string            `db:"name" json:"name"`
	Description   *string            `db:"description" json:"description"`
	CreatedAt     pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	BookmarkCount *int32             `db:"bookmark_count" json:"bookmark_count"`
	UserCount     *int32             `db:"user_count" json:"user_count"`
	Categories    []*Category        `db:"categories" json:"categories"`
	Users         []*FindUserByIdRow `db:"users" json:"users"`
}
