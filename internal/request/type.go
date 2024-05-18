package request

type RefreshAccessTokenParam struct {
	RefreshToken string `json:"refresh_token"`
}

type CreateInviteCodeParam struct {
	WorkspaceID int `json:"workspace_id"`
}

type UpdateCommentParam struct {
	CommentContext string `json:"comment_context"`
}

type CreateCommentParam struct {
	BookmarkID int    `json:"bookmark_id"`
	Context    string `json:"context"`
}

type CreateBookmarkParam struct {
	Link        string `db:"link" json:"link"`
	WorkspaceID int64  `db:"workspace_id" json:"workspace_id"`
	CategoryID  int64  `db:"category_id" json:"category_id"`
	Summary     string `db:"summary" json:"summary"`
	Title       string `db:"title" json:"title"`
}
