package bookmark

import (
	"net/http"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/request"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type BookmarkController struct {
	BookmarkService interfaces.IBookmarkService
}

// CreateBookmarkController godoc
//
//	@summary	북마크 생성
//	@schemes
//	@description	북마크를 생성합니다.
//	@tags			bookmark
//	@accept			json
//	@produce		json
//	@param			body	body		request.CreateBookmarkParam	true	"Bookmark param"
//	@success		200		{object}	repository.Bookmark
//	@failure		401		{object}	customerror.CustomError
//	@failure		404		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/bookmark [POST]
func (b *BookmarkController) CreateBookmarkController(ctx echo.Context) error {
	var body request.CreateBookmarkParam
	err := ctx.Bind(&body)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	user, err := utils.GetAuthUser(ctx)
	if err != nil {
		return err
	}
	param := repository.CreateBookmarkParams{
		Summary:     &body.Summary,
		Link:        &body.Link,
		WorkspaceID: &body.WorkspaceID,
		CategoryID:  &body.CategoryID,
		UserID:      &user.ID,
		Title:       &body.Title,
	}

	bookmark, err := b.BookmarkService.Create(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, bookmark)
}

// FindBookmarkController godoc
//
//	@summary	북마크 조회
//	@schemes
//	@description	북마크를 조회합니다.
//	@tags			bookmark
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"Bookmark id"
//	@success		200	{object}	repository.FindBookmarkRow
//	@failure		401	{object}	customerror.CustomError
//	@failure		404	{object}	customerror.CustomError
//	@failure		500	{object}	customerror.CustomError
//	@router			/api/v1/bookmark/:id [GET]
func (b *BookmarkController) FindBookmarkController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	bookmark, err := b.BookmarkService.FindById(*id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, bookmark)
}

// FindBookmarkCommentsController godoc
//
//	@summary	북마크 댓글 조회
//	@schemes
//	@description	북마크의 댓글을 조회합니다.
//	@tags			bookmark
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"Bookmark id"
//	@success		200	{object}	[]repository.Comment
//	@failure		401	{object}	customerror.CustomError
//	@failure		404	{object}	customerror.CustomError
//	@failure		500	{object}	customerror.CustomError
//	@router			/api/v1/bookmark/:id/comments [GET]
func (b *BookmarkController) FindBookmarkCommentsController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	comments, _ := b.BookmarkService.FindComments(*id)

	return ctx.JSON(http.StatusOK, comments)
}

// UpdateBookmarkController godoc
//
//	@summary	북마크 업데이트
//	@schemes
//	@description	북마크를 업데이트 합니다.
//	@tags			bookmark
//	@accept			json
//	@produce		json
//	@param			body	body		repository.UpdateBookmarkParams	true	"Bookmark param"
//	@success		200		{object}	repository.Bookmark
//	@failure		401		{object}	customerror.CustomError
//	@failure		404		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/bookmark/:id [PUT]
func (b *BookmarkController) UpdateBookmarkController(ctx echo.Context) error {
	var param repository.UpdateBookmarkParams
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	param.ID = int64(*id)
	bookmark, err := b.BookmarkService.Update(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, bookmark)
}

// FindBookmarkController godoc
//
//	@summary	북마크 삭제
//	@schemes
//	@description	북마크를 삭제합니다.
//	@tags			bookmark
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"Bookmark id"
//	@success		200	{object}	responses.OkResponse
//	@failure		401	{object}	customerror.CustomError
//	@failure		404	{object}	customerror.CustomError
//	@failure		500	{object}	customerror.CustomError
//	@router			/api/v1/bookmark/:id [DELETE]
func (b *BookmarkController) DeleteBookmarkController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	err = b.BookmarkService.Delete(*id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, responses.OkResponse{Ok: true})
}

func InitBookmarkController() *BookmarkController {
	return &BookmarkController{}
}

func (b BookmarkController) WithBookmarkService(service interfaces.IBookmarkService) BookmarkController {
	b.BookmarkService = service
	return b
}
