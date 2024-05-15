package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name ICommentService
type ICommentService interface {
	FindById(int) (*repository.Comment, error)
	Create(repository.CreateCommentParams) (*repository.Comment, error)
	Update(int, repository.UpdateCommentParams) (*repository.Comment, error)
	Delete(int, int) error
}
