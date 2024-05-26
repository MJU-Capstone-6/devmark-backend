package workspacecode

import (
	"context"
	"math/rand"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/request"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type WorkspaceCodeService struct {
	Repository       interfaces.IRepository
	CategoryService  interfaces.ICategoryService
	BookmarkService  interfaces.IBookmarkService
	WorkspaceService interfaces.IWorkspaceService
}

const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (w *WorkspaceCodeService) CreateCode(length int) *string {
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = CHARSET[rand.Intn(len(CHARSET))]
	}
	generatedString := string(randomString)
	return &generatedString
}

func (w *WorkspaceCodeService) FindByCode(code string) (*repository.FindWorkspaceCodeRow, error) {
	workspaceCode, err := w.Repository.FindWorkspaceCode(context.Background(), &code)
	if err != nil {
		return nil, customerror.WorkspaceCodeNotFound(err)
	}
	return &workspaceCode, nil
}

func (w *WorkspaceCodeService) PredictCategory(param request.PredictCategoryParam) (*repository.Bookmark, error) {
	workspaceCode, err := w.FindByCode(param.Code)
	if err != nil {
		return nil, err
	}

	dto, err := utils.PredictCategoryRequest(param.Link, param.Domain)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	category, err := w.CategoryService.FindByName(dto.Category)
	if err != nil {
		category, err = w.CategoryService.Create(dto.Category)
		if err != nil {
			return nil, err
		}
	}
	createBookmarkParam := repository.CreateBookmarkParams{
		Title:       &dto.Title,
		Link:        &param.Link,
		WorkspaceID: &workspaceCode.Workspace.ID,
		CategoryID:  &category.ID,
		UserID:      &param.UserID,
	}

	bookmark, err := w.BookmarkService.Create(createBookmarkParam)
	if err != nil {
		return nil, err
	}
	registerParam := repository.RegisterCategoryToWorkspaceParams{
		WorkspaceID: workspaceCode.Workspace.ID,
		CategoryID:  category.ID,
	}
	err = w.WorkspaceService.RegisterCategory(registerParam)
	if err != nil {
		return nil, err
	}

	return bookmark, nil
}

func (w *WorkspaceCodeService) Update(param repository.UpdateWorkspaceCodeParams) (*repository.WorkspaceCode, error) {
	workspaceCode, err := w.Repository.UpdateWorkspaceCode(context.Background(), param)
	if err != nil {
		return nil, customerror.WorkspaceCodeUpdateFail(err)
	}
	return &workspaceCode, nil
}

func (w *WorkspaceCodeService) Create(param repository.CreateWorkspaceCodeParams) (*repository.WorkspaceCode, error) {
	code := w.CreateCode(6)
	param.Code = code
	workspaceCode, err := w.Repository.CreateWorkspaceCode(context.Background(), param)
	if err != nil {
		return nil, customerror.WorkspaceCodeCreationFail(err)
	}
	return &workspaceCode, nil
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

func (w WorkspaceCodeService) WithWorkspaceService(service interfaces.IWorkspaceService) WorkspaceCodeService {
	w.WorkspaceService = service
	return w
}
