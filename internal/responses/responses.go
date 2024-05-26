package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetKakaoInfoResponse struct {
	AccessToken  string `json:"access_key"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type OkResponse struct {
	Ok bool `json:"ok"`
}

type FindWorkspaceResponse struct {
	WorkspaceName string `json:"workspace_name"`
}

func Unauthorized(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusUnauthorized, data)
}

func NotAcceptable(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusNotAcceptable, data)
}

func NotFound(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusNotFound, data)
}

func BadRequest(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusBadRequest, data)
}

func OK(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, data)
}

func InternalServer(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusInternalServerError, data)
}
