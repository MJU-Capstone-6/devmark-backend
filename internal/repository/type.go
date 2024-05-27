package repository

import "github.com/jackc/pgx/v5/pgtype"

type BookmarkCommentRow struct {
	UserID         *int64             `db:"user_id" json:"user_id"`
	Username       *string            `db:"username" json:"username"`
	CommentID      *int64             `db:"comment_id" json:"comment_id"`
	CreatedAt      pgtype.Timestamptz `db:"created_at" json:"created_at"`
	CommentContext *string            `db:"comment_context" json:"comment_context"`
}

type UserBookmarkCount struct {
	UserID        *int64  `db:"user_id" json:"user_id"`
	Username      *string `db:"username" json:"username"`
	BookmarkCount *int64  `db:"bookmark_count" json:"bookmark_count"`
}
