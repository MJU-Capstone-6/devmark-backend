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
