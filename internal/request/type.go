package request

type RefreshAccessTokenParam struct {
	RefreshToken string `json:"refresh_token"`
}

type CreateInviteCodeParam struct {
	WorkspaceID int `json:"workspace_id"`
}
