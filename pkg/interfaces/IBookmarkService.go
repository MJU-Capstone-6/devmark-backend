package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

type IBookmarkService interface {
	Create(repository.CreateBookmarkParams) (*repository.Bookmark, error)
	Update(repository.UpdateBookmarkParams) (*repository.Bookmark, error)
	FindById(int) (*repository.FindBookmarkRow, error)
	Delete(int) error
}
