package workspace

import (
	"context"
	"errors"
	"log"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type WorkspaceService struct {
	Repository        interfaces.IRepository
	InviteCodeService interfaces.IInviteCodeService
}

func (w *WorkspaceService) Create(userId int, createWorkspaceParam repository.CreateWorkspaceParams) (*repository.Workspace, error) {
	workspace, err := w.Repository.CreateWorkspace(context.Background(), createWorkspaceParam)
	if err != nil {
		return nil, customerror.WorkspaceCreateFail(err)
	}
	param := repository.JoinWorkspaceWithoutCodeParams{
		WorkspaceID: workspace.ID,
		UserID:      int64(userId),
	}
	err = w.Repository.JoinWorkspaceWithoutCode(context.Background(), param)
	if err != nil {
		return nil, customerror.WorkspaceJoinFailErr(err)
	}
	return &workspace, nil
}

func (w *WorkspaceService) FindById(id int) (*repository.FindWorkspaceRow, error) {
	workspaceId := int64(id)
	_, err := w.Repository.CheckWorkspaceExists(context.Background(), workspaceId)
	if err != nil {
		return nil, customerror.WorkspaceNotFoundErr(err)
	}

	workspace, err := w.Repository.FindWorkspace(context.Background(), workspaceId)
	if err != nil {
		return &repository.FindWorkspaceRow{ID: workspaceId, Categories: []*repository.Category{}, Users: []*repository.FindUserByIdRow{}}, nil
	}
	return &workspace, nil
}

func (w *WorkspaceService) Update(param repository.UpdateWorkspaceParams) (*repository.Workspace, error) {
	workspace, err := w.Repository.UpdateWorkspace(context.Background(), param)
	if err != nil {
		return nil, customerror.WorkspaceUpdateFail(err)
	}
	return &workspace, nil
}

func (w *WorkspaceService) Delete(id int) error {
	err := w.Repository.DeleteWorkspace(context.Background(), int64(id))
	if err != nil {
		return customerror.WorkspaceDeleteFail(err)
	}
	return nil
}

func (w *WorkspaceService) Join(code string, param repository.JoinWorkspaceParams) error {
	inviteCode, err := w.InviteCodeService.FindByCode(code)
	if err != nil {
		log.Println(err)
		return err
	}

	verifyParam := utils.VerifyCodeParam{
		Code:        code,
		WorkspaceId: int(*inviteCode.WorkspaceID),
	}
	if ok, err := w.InviteCodeService.VerifyCode(verifyParam); !ok {
		if err != nil {
			return err
		}
		return customerror.CodeVerifyFailedErr(nil)
	}

	param.WorkspaceID = int64(*inviteCode.WorkspaceID)
	findJoinedUserParam := repository.FindWorkspaceJoinedUserParams(param)
	_, err = w.Repository.FindWorkspaceJoinedUser(context.Background(), findJoinedUserParam)

	if err == nil {
		return customerror.WorkspaceAlreadyJoined(errors.New(""))
	}

	err = w.Repository.JoinWorkspace(context.Background(), param)
	if err != nil {
		return customerror.WorkspaceJoinFailErr(err)
	}
	return nil
}

func (w *WorkspaceService) FindCategoriesById(id int) (*[]*repository.Category, error) {
	categories, err := w.Repository.FindWorkspaceCategory(context.Background(), int64(id))
	if err != nil {
		return &[]*repository.Category{}, nil
	}
	return &categories, nil
}

func (w *WorkspaceService) FindCategoryBookmark(param repository.FindWorkspaceCategoryBookmarkParams) (*[]repository.Bookmark, error) {
	bookmarks, err := w.Repository.FindWorkspaceCategoryBookmark(context.Background(), param)
	if err != nil {
		return &[]repository.Bookmark{}, nil
	}
	return &bookmarks, nil
}

func (w *WorkspaceService) SearchBookmark(param repository.SearchWorkspaceBookmarkParams) (*[]repository.SearchWorkspaceBookmarkRow, error) {
	bookmarks, err := w.Repository.SearchWorkspaceBookmark(context.Background(), param)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	if bookmarks != nil {
		return &bookmarks, nil
	} else {
		return &[]repository.SearchWorkspaceBookmarkRow{}, nil
	}
}

func (w *WorkspaceService) RegisterCategory(param repository.RegisterCategoryToWorkspaceParams) error {
	err := w.Repository.RegisterCategoryToWorkspace(context.Background(), param)
	if err != nil {
		return customerror.WorkspaceRegisterCategoryFail(err)
	}
	return nil
}

func (w *WorkspaceService) FindInfoById(id int) (*repository.FindWorkspaceInfoRow, error) {
	workspaceInfo, err := w.Repository.FindWorkspaceInfo(context.Background(), int64(id))
	if err != nil {
		return nil, customerror.WorkspaceNotFoundErr(err)
	}
	return &workspaceInfo, nil
}

func (w *WorkspaceService) FindTopRecommendLinks(id int) (*[]repository.FindRecommendLinksRow, error) {
	parsedID := int64(id)
	recommendLinks, err := w.Repository.FindRecommendLinks(context.Background(), &parsedID)
	if err != nil {
		return &[]repository.FindRecommendLinksRow{}, nil
	}
	return &recommendLinks, nil
}

func (w *WorkspaceService) FindTopCategories(id int) (*[]repository.FindTopCategoriesRow, error) {
	parsedID := int64(id)
	categories, err := w.Repository.FindTopCategories(context.Background(), &parsedID)
	if err != nil {
		return &[]repository.FindTopCategoriesRow{}, nil
	}
	return &categories, nil
}

func (w *WorkspaceService) Exit(param repository.ExitWorkspaceParams) error {
	_, err := w.FindById(int(param.WorkspaceID))
	if err != nil {
		return err
	}
	_, err = w.Repository.IsUserJoinedWorkspace(context.Background(), repository.IsUserJoinedWorkspaceParams(param))
	if err != nil {
		log.Println(err)
		return customerror.WorkspaceNeverJoined(err)
	}
	err = w.Repository.ExitWorkspace(context.Background(), param)
	if err != nil {
		return customerror.WorkspaceExitFail(err)
	}
	return nil
}

func InitWorkspaceService(repo interfaces.IRepository) *WorkspaceService {
	return &WorkspaceService{Repository: repo}
}

func (w WorkspaceService) WithInviteCodeService(service interfaces.IInviteCodeService) WorkspaceService {
	w.InviteCodeService = service
	return w
}
