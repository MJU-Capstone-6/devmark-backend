package bookmark

import (
	"context"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type BookmarkService struct {
	Repository interfaces.IRepository
}

func (b *BookmarkService) Create(param repository.CreateBookmarkParams) (*repository.Bookmark, error) {
	bookmark, err := b.Repository.CreateBookmark(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return &bookmark, nil
}

func (b *BookmarkService) Update(param repository.UpdateBookmarkParams) (*repository.Bookmark, error) {
	bookmark, err := b.Repository.UpdateBookmark(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return &bookmark, nil
}

func (b *BookmarkService) FindById(id int) (*repository.FindBookmarkRow, error) {
	bookmark, err := b.Repository.FindBookmark(context.Background(), int64(id))
	if err != nil {
		return nil, err
	}
	return &bookmark, nil
}

func (b *BookmarkService) Delete(id int) error {
	err := b.Repository.DeleteBookmark(context.Background(), int64(id))
	if err != nil {
		return err
	}
	return nil
}

func InitBookmarkService() *BookmarkService {
	return &BookmarkService{}
}

func (b BookmarkService) WithRepository(repo interfaces.IRepository) BookmarkService {
	b.Repository = repo
	return b
}
