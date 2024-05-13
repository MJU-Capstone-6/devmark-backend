package workspace

type CreateWorkspaceParam struct {
	Name string `json:"name"`
}

type JoinWorkspaceParam struct {
	Code string `json:"code"`
}
