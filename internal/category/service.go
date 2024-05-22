package category

import (
	"context"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type CategoryService struct {
	Repository interfaces.IRepository
}

func (c *CategoryService) FindById(id int) (*repository.Category, error) {
	category, err := c.Repository.FindCategoryById(context.Background(), int64(id))
	if err != nil {
		return nil, customerror.CategoryNotFound(err)
	}
	return &category, nil
}

func (c *CategoryService) FindByName(name string) (*repository.Category, error) {
	category, err := c.Repository.FindCategoryByName(context.Background(), &name)
	if err != nil {
		return nil, customerror.CategoryNotFound(err)
	}
	return &category, nil
}

func (c *CategoryService) Create(name string) (*repository.Category, error) {
	_, err := c.FindByName(name)
	if err != nil {
		return nil, customerror.CategoryAlreadyExists(err)
	} else {
		newCategory, err := c.Repository.CreateCategory(context.Background(), &name)
		if err != nil {
			return nil, customerror.CategoryCreationFail(err)
		}
		return &newCategory, nil
	}
}

func (c *CategoryService) Update(param repository.UpdateCategoryParams) (*repository.Category, error) {
	updatedCategory, err := c.Repository.UpdateCategory(context.Background(), param)
	if err != nil {
		return nil, customerror.CategoryUpdateFail(err)
	}
	return &updatedCategory, nil
}

func (c *CategoryService) Delete(id int) error {
	err := c.Repository.DeleteCategory(context.Background(), int64(id))
	if err != nil {
		return customerror.CategoryDeleteFail(err)
	}
	return nil
}

func InitCategoryService() *CategoryService {
	return &CategoryService{}
}

func (c CategoryService) WithRepository(repo interfaces.IRepository) CategoryService {
	c.Repository = repo
	return c
}
