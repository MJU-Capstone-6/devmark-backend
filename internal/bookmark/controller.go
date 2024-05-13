package bookmark

import (
	"net/http"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type BookmarkController struct {
	BookmarkService interfaces.IBookmarkService
}

func (b *BookmarkController) CreateBookmarkController(ctx echo.Context) error {
	var body repository.CreateBookmarkParams
	err := ctx.Bind(&body)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	bookmark, err := b.BookmarkService.Create(body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, bookmark)
}

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

func (b *BookmarkController) DeleteBookmarkController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	err = b.BookmarkService.Delete(*id)
	if err != nil {
		return err
	}
	return nil
}

func InitBookmarkController() *BookmarkController {
	return &BookmarkController{}
}

func (b BookmarkController) WithBookmarkService(service interfaces.IBookmarkService) BookmarkController {
	b.BookmarkService = service
	return b
}
