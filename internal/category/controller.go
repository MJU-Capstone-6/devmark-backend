package category

import (
	"net/http"

	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryService interfaces.ICategoryService
}

// CreateCategoryController godoc
//
//	@summary	Create Category
//	@schemes
//	@description	카테고리를 생성합니다. *워크스페이스에 카테고리를 등록하는 API는 따로 존재합니다.
//	@tags			category
//	@accept			json
//	@produce		json
//	@param			body	body		CreateCategoryParam	true	"body to create category"
//	@success		200		{object}	repository.Category
//	@failure		400		{object}	customerror.CustomError
//	@failure		401		{object}	customerror.CustomError
//	@failure		422		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/category [POST]
func (c *CategoryController) CreateCategoryController(ctx echo.Context) error {
	var param CreateCategoryParam
	err := ctx.Bind(&param)
	if err != nil {
		return err
	}
	category, err := c.CategoryService.Create(param.Name)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, category)
}

// UpdateCategoryController godoc
//
//	@summary	Update Category
//	@schemes
//	@description	카테고리를 업데이트 합니다.
//	@tags			category
//	@accept			json
//	@produce		json
//	@param			id		path		int								true	"category id"
//	@param			body	body		repository.UpdateCategoryParams	true	"body to update category"
//	@success		200		{object}	repository.Category
//	@failure		400		{object}	customerror.CustomError
//	@failure		401		{object}	customerror.CustomError
//	@failure		422		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/category/:id [PUT]
func (c *CategoryController) UpdateCategoryController(ctx echo.Context) error {
	var param repository.UpdateCategoryParams
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	err = ctx.Bind(&param)
	if err != nil {
		return err
	}
	param.ID = int64(*id)
	category, err := c.CategoryService.Update(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, category)
}

// DeleteCategoryController godoc
//
//	@summary	Delete Category
//	@schemes
//	@description	카테고리를 삭제 합니다.
//	@tags			category
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"category id"
//	@success		200	{object}	repository.Category
//	@failure		400	{object}	customerror.CustomError
//	@failure		401	{object}	customerror.CustomError
//	@failure		422	{object}	customerror.CustomError
//	@failure		500	{object}	customerror.CustomError
//	@router			/api/v1/category/:id [DELETE]
func (c *CategoryController) DeleteCategoryController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	err = c.CategoryService.Delete(*id)
	if err != nil {
		return err
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
