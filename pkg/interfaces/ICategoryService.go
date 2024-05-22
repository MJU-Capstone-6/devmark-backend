package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name ICategoryService
type ICategoryService interface {
	FindById(int) (*repository.Category, error)
	FindByName(string) (*repository.Category, error)
	Create(string) (*repository.Category, error)
	Update(repository.UpdateCategoryParams) (*repository.Category, error)
	Delete(int) error
}
