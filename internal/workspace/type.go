package workspace

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

type CreateWorkspaceParam struct {
	Name string `json:"name"`
}

type JoinWorkspaceParam struct {
	Code string `json:"code"`
	repository.JoinWorkspaceParams
}
