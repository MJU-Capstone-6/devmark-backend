package repository

type BookmarkCommentRow struct {
	UserID         *int64  `db:"user_id" json:"user_id"`
	Username       *string `db:"username" json:"username"`
	CommentID      *int64  `db:"comment_id" json:"comment_id"`
	CommentContext *string `db:"comment_context" json:"comment_context"`
}
