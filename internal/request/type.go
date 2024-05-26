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

type PredictCategoryParam struct {
	Code   string `json:"code"`
	Link   string `json:"link"`
	Domain string `json:"domain"`
	UserID int64  `json:"user_id"`
}

type PredictCategoryBody struct {
	Code string `json:"code"`
	Link string `json:"link"`
}

type FindWorkspaceCodeParam struct {
	Code string `json:"code"`
}

type UpdateBookmarkParam struct {
	Link         *string `db:"link" json:"link"`
	WorkspaceID  *int64  `db:"workspace_id" json:"workspace_id"`
	CategoryName *string `db:"category_name" json:"category_name"`
	Summary      *string `db:"summary" json:"summary"`
	Title        *string `db:"title" json:"title"`
}
