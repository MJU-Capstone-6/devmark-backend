package bookmark

import (
	"context"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/request"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type BookmarkService struct {
	Repository       interfaces.IRepository
	WorkspaceService interfaces.IWorkspaceService
	CategoryService  interfaces.ICategoryService
}

func (b *BookmarkService) Create(param repository.CreateBookmarkParams) (*repository.Bookmark, error) {
	duplicateParam := repository.FindDuplicateBookmarkParams{
		WorkspaceID: param.WorkspaceID,
		Link:        param.Link,
	}
	_, err := b.Repository.FindDuplicateBookmark(context.Background(), duplicateParam)
	if err == nil {
		return nil, customerror.BookmarkDuplicated(err)
	}
	_, err = b.WorkspaceService.FindById(int(*param.WorkspaceID))
	if err != nil {
		return nil, err
	}
	_, err = b.CategoryService.FindById(int(*param.CategoryID))
	if err != nil {
		return nil, err
	}

	bookmark, err := b.Repository.CreateBookmark(context.Background(), param)
	if err != nil {
		return nil, customerror.BookmarkCreationFail(err)
	}
	return &bookmark, nil
}

func (b *BookmarkService) Update(id int, param request.UpdateBookmarkParam) (*repository.Bookmark, error) {
	var category *repository.Category
	category, err := b.CategoryService.FindByName(*param.CategoryName)
	if err != nil {
		category, err = b.CategoryService.Create(*param.CategoryName)
		if err != nil {
			return nil, err
		}
	}
	updateParam := repository.UpdateBookmarkParams{
		ID:          int64(id),
		WorkspaceID: param.WorkspaceID,
		CategoryID:  &category.ID,
		Summary:     param.Summary,
		Title:       param.Title,
		Link:        param.Link,
	}
	bookmark, err := b.Repository.UpdateBookmark(context.Background(), updateParam)
	if err != nil {
		return nil, customerror.BookmarkUpdateFail(err)
	}
	return &bookmark, nil
}

func (b *BookmarkService) FindById(id int) (*repository.FindBookmarkRow, error) {
	bookmark, err := b.Repository.FindBookmark(context.Background(), int64(id))
	if err != nil {
		return nil, customerror.BookmarkNotFound(err)
	}
	return &bookmark, nil
}

func (b *BookmarkService) ReadBookmark(id int) error {
	err := b.Repository.ReadBookmark(context.Background(), int64(id))
	if err != nil {
		return customerror.BookmarkUpdateFail(err)
	}
	return nil
}

func (b *BookmarkService) FindComments(id int) (*[]*repository.BookmarkCommentRow, error) {
	comments, err := b.Repository.FindBookmarkComment(context.Background(), int64(id))
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	if comments != nil {
		return &comments, nil
	} else {
		return &[]*repository.BookmarkCommentRow{}, nil
	}
}

func (b *BookmarkService) Delete(id int) error {
	err := b.Repository.DeleteBookmark(context.Background(), int64(id))
	if err != nil {
		return customerror.BookmarkDeleteFail(err)
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

func (b BookmarkService) WithWorkspaceService(service interfaces.IWorkspaceService) BookmarkService {
	b.WorkspaceService = service
	return b
}

func (b BookmarkService) WithCategoryService(service interfaces.ICategoryService) BookmarkService {
	b.CategoryService = service
	return b
}
