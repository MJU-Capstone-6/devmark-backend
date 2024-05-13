package category

import (
	"net/http"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryService interfaces.ICategoryService
}

func (c *CategoryController) CreateCategoryController(ctx echo.Context) error {
	var param CreateCategoryParam
	err := ctx.Bind(&param)
	if err != nil {
		return responses.BadRequest(ctx, customerror.CategoryBodyNotProvide(err))
	}
	category, err := c.CategoryService.Create(param.Name)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	return ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) UpdateCategoryController(ctx echo.Context) error {
	var param repository.UpdateCategoryParams
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	err = ctx.Bind(&param)
	if err != nil {
		return responses.BadRequest(ctx, customerror.CategoryBodyNotProvide(err))
	}
	param.ID = int64(*id)
	category, err := c.CategoryService.Update(param)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}
	return ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) DeleteCategoryController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	err = c.CategoryService.Delete(*id)
	if err != nil {
		return responses.InternalServer(ctx, customerror.InternalServerError(err))
	}

	return nil
}

func InitCategoryController() *CategoryController {
	return &CategoryController{}
}

func (c CategoryController) WithCategoryService(service interfaces.ICategoryService) CategoryController {
	c.CategoryService = service
	return c
}
