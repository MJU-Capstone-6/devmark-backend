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
}

type Category struct {
	ID        int64              `db:"id" json:"id"`
	Name      *string            `db:"name" json:"name"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

type Comment struct {
	ID         int64              `db:"id" json:"id"`
	BookmarkID *int64             `db:"bookmark_id" json:"bookmark_id"`
	UserID     *int64             `db:"user_id" json:"user_id"`
	Context    *int64             `db:"context" json:"context"`
	CreatedAt  pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
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
	ID         *int64      `db:"id" json:"id"`
	Workspaces []Workspace `db:"workspaces" json:"workspaces"`
}

type Workspace struct {
	ID        int64              `db:"id" json:"id"`
	Name      *string            `db:"name" json:"name"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

type WorkspaceCategory struct {
	WorkspaceID int64 `db:"workspace_id" json:"workspace_id"`
	CategoryID  int64 `db:"category_id" json:"category_id"`
}

type WorkspaceUser struct {
	WorkspaceID int64 `db:"workspace_id" json:"workspace_id"`
	UserID      int64 `db:"user_id" json:"user_id"`
}

type WorkspaceUserCategory struct {
	ID         int64              `db:"id" json:"id"`
	Name       *string            `db:"name" json:"name"`
	CreatedAt  pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	Categories []Category         `db:"categories" json:"categories"`
	Users      []User             `db:"users" json:"users"`
}
