package app

import (
	"fmt"

	"github.com/MJU-Capstone-6/devmark-backend/internal/auth"
	"github.com/MJU-Capstone-6/devmark-backend/internal/bookmark"
	"github.com/MJU-Capstone-6/devmark-backend/internal/category"
	invitecode "github.com/MJU-Capstone-6/devmark-backend/internal/inviteCode"
	"github.com/MJU-Capstone-6/devmark-backend/internal/jwtToken"
	"github.com/MJU-Capstone-6/devmark-backend/internal/middlewares"
	refreshtoken "github.com/MJU-Capstone-6/devmark-backend/internal/refreshToken"
	"github.com/MJU-Capstone-6/devmark-backend/internal/user"
	"github.com/MJU-Capstone-6/devmark-backend/internal/workspace"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const V1 = "/api/v1"

func (app *Application) InitRoutes() {
	app.Handler.HTTPErrorHandler = CustomHTTPErrorHandler
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Handler.File("/docs", "swagger.json")
	app.Handler.GET("/api/v1/swagger/*", echoSwagger.WrapHandler)
	app.InitUserRoutes()
	app.InitAuthRoutes()
	app.InitWorkspaceRoutes()
	app.InitInviteCodeRoutes()
	app.InitCategoryRoutes()
	app.InitBookmarkRoutes()
}

func (app *Application) InitUserRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/user", V1))

	userService := user.InitUserService(&app.Repository)
	userController := user.InitController().WithUserService(userService)
	e.GET("/:name", userController.ViewOneUser)
}

func (app *Application) InitAuthRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/auth", V1))
	userService := user.InitUserService(&app.Repository)
	jwtService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)
	refreshTokenService := refreshtoken.InitRefreshTokenService(&app.Repository, jwtService)
	authService := auth.InitAuthService(&app.Repository, userService, jwtService, refreshTokenService)
	authController := auth.InitAuthController().
		WithKakaoInfo(app.Config.Kakao).
		WithAuthService(authService)
	customMiddleware := middlewares.InitMiddleware()
	e.POST("/:provider", authController.GetKakaoUserInfo, customMiddleware.ParseHeader)
}

func (app *Application) InitWorkspaceRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/workspace", V1))
	tempWorkspaceService := workspace.InitWorkspaceService(&app.Repository)
	inviteCodeService := invitecode.InitInviteCodeService().WithRepository(&app.Repository).WithWorkspaceService(tempWorkspaceService)
	workspaceService := workspace.InitWorkspaceService(&app.Repository).WithInviteCodeService(&inviteCodeService)
	workspaceController := workspace.InitWorkspaceController().WithWorkspaceService(&workspaceService)

	userService := user.InitUserService(&app.Repository)
	jwtService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)

	customMiddleware := middlewares.InitMiddleware().WithUserService(userService).WithJwtTokenService(jwtService)

	e.GET("/:id", workspaceController.ViewWorkspaceController, customMiddleware.Auth)
	e.PUT("/:id", workspaceController.UpdateWorkspaceController, customMiddleware.Auth)
	e.POST("", workspaceController.CreateWorkspaceController, customMiddleware.Auth)
	e.POST("/:id/join", workspaceController.JoinWorkspaceController, customMiddleware.Auth)
	e.DELETE("/:id", workspaceController.DeleteWorkspaceController, customMiddleware.Auth)
}

func (app *Application) InitInviteCodeRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/invitecode", V1))
	workspaceService := workspace.InitWorkspaceService(&app.Repository)
	inviteCodeService := invitecode.InitInviteCodeService().
		WithRepository(&app.Repository).
		WithWorkspaceService(workspaceService)
	inviteCodeController := invitecode.InitInviteCodeController().WithInviteCodeService(&inviteCodeService)
	jwtService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)
	userService := user.InitUserService(&app.Repository)

	customMiddleware := middlewares.InitMiddleware().WithUserService(userService).WithJwtTokenService(jwtService)
	e.POST("", inviteCodeController.GenerateInviteCodeController, customMiddleware.Auth)
}

func (app *Application) InitCategoryRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/category", V1))

	categoryService := category.InitCategoryService().WithRepository(&app.Repository)
	categoryController := category.InitCategoryController().WithCategoryService(&categoryService)

	jwtService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)
	userService := user.InitUserService(&app.Repository)

	customMiddleware := middlewares.InitMiddleware().WithUserService(userService).WithJwtTokenService(jwtService)

	e.POST("", categoryController.CreateCategoryController, customMiddleware.Auth)
	e.PUT("/:id", categoryController.UpdateCategoryController, customMiddleware.Auth)
	e.DELETE("/:id", categoryController.DeleteCategoryController, customMiddleware.Auth)
}

func (app *Application) InitBookmarkRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/bookmark", V1))
	workspaceService := workspace.InitWorkspaceService(&app.Repository)
	categoryService := category.InitCategoryService().WithRepository(&app.Repository)
	bookmarkService := bookmark.InitBookmarkService().
		WithRepository(&app.Repository).
		WithWorkspaceService(workspaceService).
		WithCategoryService(&categoryService)
	bookmarkController := bookmark.InitBookmarkController().WithBookmarkService(&bookmarkService)

	e.GET("/:id", bookmarkController.FindBookmarkController)
	e.POST("", bookmarkController.CreateBookmarkController)
	e.PUT("/:id", bookmarkController.UpdateBookmarkController)
	e.DELETE("/:id", bookmarkController.DeleteBookmarkController)
}
