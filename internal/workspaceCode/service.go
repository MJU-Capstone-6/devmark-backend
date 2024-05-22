package workspacecode

import (
	"context"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type WorkspaceCodeService struct {
	Repository      interfaces.IRepository
	CategoryService interfaces.ICategoryService
	BookmarkService interfaces.IBookmarkService
}

func (w *WorkspaceCodeService) FindByCode(code string) (*repository.FindWorkspaceCodeRow, error) {
	workspaceCode, err := w.Repository.FindWorkspaceCode(context.Background(), &code)
	if err != nil {
		return nil, customerror.WorkspaceCodeNotFound(err)
	}
	return &workspaceCode, nil
}

func (w *WorkspaceCodeService) CreateCategoryAndBookmark(code string) (*repository.Bookmark, error) {
	return nil, nil
}

func (w *WorkspaceCodeService) Update(param repository.UpdateWorkspaceCodeParams) (*repository.WorkspaceCode, error) {
	return nil, nil
}

func (w *WorkspaceCodeService) Create(param repository.CreateWorkspaceCodeParams) (*repository.WorkspaceCode, error) {
	return nil, nil
}

func InitWorkspaceCodeService() *WorkspaceCodeService {
	return &WorkspaceCodeService{}
}

func (w WorkspaceCodeService) WithRepository(repo interfaces.IRepository) WorkspaceCodeService {
	w.Repository = repo
	return w
}

func (w WorkspaceCodeService) WithCategoryService(service interfaces.ICategoryService) WorkspaceCodeService {
	w.CategoryService = service
	return w
}

func (w WorkspaceCodeService) WithBookmarkService(service interfaces.IBookmarkService) WorkspaceCodeService {
	w.BookmarkService = service
	return w
}
