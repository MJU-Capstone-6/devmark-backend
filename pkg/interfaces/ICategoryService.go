package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name ICategoryService
type ICategoryService interface {
	Create(string) (*repository.Category, error)
	Update(repository.UpdateCategoryParams) (*repository.Category, error)
	Delete(int) error
}
