package interfaces

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/request"
)

type IBookmarkService interface {
	Create(repository.CreateBookmarkParams) (*repository.Bookmark, error)
	Update(int, request.UpdateBookmarkParam) (*repository.Bookmark, error)
	FindById(int) (*repository.FindBookmarkRow, error)
	FindComments(int) (*[]*repository.BookmarkCommentRow, error)
	Delete(int) error
	ReadBookmark(int) error
}
