package comment

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

type CommentController struct {
	CommentService interfaces.ICommentService
}

// CreateCommentController godoc
//
//	@summary	댓글을 생성합니다.
//	@schemes
//	@description	댓글을 생성합니다.
//	@tags			comment
//	@accept			json
//	@produce		json
//	@param			param	body		request.CreateCommentParam	true	"comment param"
//	@success		200		{object}	repository.Comment
//	@failure		400		{object}	customerror.CustomError
//	@failure		401		{object}	customerror.CustomError
//	@failure		422		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/comment [POST]
func (c *CommentController) CreateCommentController(ctx echo.Context) error {
	var requestParam request.CreateCommentParam
	err := ctx.Bind(&requestParam)
	if err != nil {
		return customerror.CommentParamNotValid(err)
	}
	user, err := utils.GetAuthUser(ctx)
	if err != nil {
		return err
	}
	bookmarkID := int64(requestParam.BookmarkID)
	param := repository.CreateCommentParams{
		BookmarkID:     &bookmarkID,
		CommentContext: &requestParam.Context,
		UserID:         &user.ID,
	}

	comment, err := c.CommentService.Create(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, comment)
}

// UpdateCommentController godoc
//
//	@summary	댓글을 업데이트 합니다.
//	@schemes
//	@description	댓글을 업데이트 합니다.
//	@tags			comment
//	@accept			json
//	@produce		json
//	@param			param	body		request.UpdateCommentParam	true	"comment param"
//	@param			id		path		int							true	"comment id"
//	@success		200		{object}	repository.Comment
//	@failure		400		{object}	customerror.CustomError
//	@failure		401		{object}	customerror.CustomError
//	@failure		422		{object}	customerror.CustomError
//	@failure		500		{object}	customerror.CustomError
//	@router			/api/v1/comment/:id [PUT]
func (c *CommentController) UpdateCommentController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	var requestParam request.UpdateCommentParam
	err = ctx.Bind(&requestParam)
	if err != nil || requestParam.CommentContext == "" {
		return customerror.CommentParamNotValid(err)
	}

	user, err := utils.GetAuthUser(ctx)
	if err != nil {
		return err
	}

	param := repository.UpdateCommentParams{
		CommentContext: &requestParam.CommentContext,
		ID:             int64(*id),
	}

	updatedComment, err := c.CommentService.Update(int(user.ID), param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, updatedComment)
}

// DeleteCommentController godoc
//
//	@summary	댓글을 삭제합니다.
//	@schemes
//	@description	댓글을 삭제합니다.
//	@tags			comment
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"comment id"
//	@success		200	{object}	responses.OkResponse
//	@failure		400	{object}	customerror.CustomError
//	@failure		401	{object}	customerror.CustomError
//	@failure		422	{object}	customerror.CustomError
//	@failure		500	{object}	customerror.CustomError
//	@router			/api/v1/comment/:id [DELETE]
func (c *CommentController) DeleteCommentController(ctx echo.Context) error {
	id, err := utils.ParseURLParam(ctx, "id")
	if err != nil {
		return err
	}
	user, err := utils.GetAuthUser(ctx)
	if err != nil {
		return err
	}

	err = c.CommentService.Delete(int(user.ID), *id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, responses.OkResponse{Ok: true})
}

func InitCommentController() *CommentController {
	return &CommentController{}
}

func (c CommentController) WithCommentService(service interfaces.ICommentService) CommentController {
	c.CommentService = service
	return c
}
