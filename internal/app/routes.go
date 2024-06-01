package app

import (
	"fmt"
	"net/http"

	"github.com/MJU-Capstone-6/devmark-backend/internal/auth"
	"github.com/MJU-Capstone-6/devmark-backend/internal/bookmark"
	"github.com/MJU-Capstone-6/devmark-backend/internal/category"
	"github.com/MJU-Capstone-6/devmark-backend/internal/comment"
	deviceinfo "github.com/MJU-Capstone-6/devmark-backend/internal/deviceInfo"
	"github.com/MJU-Capstone-6/devmark-backend/internal/gpt"
	invitecode "github.com/MJU-Capstone-6/devmark-backend/internal/inviteCode"
	"github.com/MJU-Capstone-6/devmark-backend/internal/jwtToken"
	"github.com/MJU-Capstone-6/devmark-backend/internal/middlewares"
	refreshtoken "github.com/MJU-Capstone-6/devmark-backend/internal/refreshToken"
	"github.com/MJU-Capstone-6/devmark-backend/internal/user"
	"github.com/MJU-Capstone-6/devmark-backend/internal/workspace"
	workspacecode "github.com/MJU-Capstone-6/devmark-backend/internal/workspaceCode"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const V1 = "/api/v1"

func (app *Application) InitRoutes() {
	app.Handler.HTTPErrorHandler = CustomHTTPErrorHandler
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Handler.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	app.Handler.File("/docs", "swagger.json")
	app.Handler.GET("/api/v1/swagger/*", echoSwagger.WrapHandler)
	app.InitUserRoutes()
	app.InitAuthRoutes()
	app.InitWorkspaceRoutes()
	app.InitInviteCodeRoutes()
	app.InitCategoryRoutes()
	app.InitBookmarkRoutes()
	app.InitRefreshTokenRoutes()
	app.InitCommentRoutes()
	app.InitWorkspaceCodeRoutes()
}

func (app *Application) InitUserRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/user", V1))

	userService := user.InitUserService(&app.Repository)
	userController := user.InitController().WithUserService(userService)

	jwtService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)
	customMiddleware := middlewares.InitMiddleware().WithUserService(userService).WithJwtTokenService(jwtService)
	e.GET("/workspace", userController.ViewUserWorkspace, customMiddleware.Auth)
}

func (app *Application) InitAuthRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/auth", V1))
	userService := user.InitUserService(&app.Repository)
	jwtService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)
	refreshTokenService := refreshtoken.InitRefreshTokenService(&app.Repository, jwtService)
	deviceInfoService := deviceinfo.InitDeviceInfoService().WithRepository(&app.Repository)
	authService := auth.InitAuthService(&app.Repository, userService, jwtService, refreshTokenService, &deviceInfoService)
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
	categoryService := category.InitCategoryService().WithRepository(&app.Repository)
	bookmarkService := bookmark.InitBookmarkService().
		WithRepository(&app.Repository).
		WithWorkspaceService(&workspaceService).
		WithCategoryService(&categoryService)

	workspaceCodeService := workspacecode.InitWorkspaceCodeService().WithRepository(&app.Repository).WithBookmarkService(&bookmarkService).WithCategoryService(&categoryService)
	workspaceController = workspaceController.WithWorkspaceCodeService(&workspaceCodeService)
	userService := user.InitUserService(&app.Repository)
	jwtService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)

	customMiddleware := middlewares.InitMiddleware().WithUserService(userService).WithJwtTokenService(jwtService)

	e.GET("/:id", workspaceController.ViewWorkspaceController, customMiddleware.Auth)
	e.GET("/:id/category", workspaceController.FindWorkspaceCategoriesController, customMiddleware.Auth)
	e.GET("/:workspace_id/category/:category_id", workspaceController.FindWorkspaceCategoryBookmark, customMiddleware.Auth)
	e.GET("/:id/bookmark", workspaceController.SearchBookmarkController, customMiddleware.Auth)
	e.GET("/:id/info", workspaceController.FindWorkspaceInfoController, customMiddleware.Auth)
	e.PUT("/:id", workspaceController.UpdateWorkspaceController, customMiddleware.Auth)
	e.POST("", workspaceController.CreateWorkspaceController, customMiddleware.Auth)
	e.POST("/join", workspaceController.JoinWorkspaceController, customMiddleware.Auth)
	e.POST("/:workspace_id/category/:category_id", workspaceController.RegisterCategoryToWorkspaceController, customMiddleware.Auth)
	e.POST("/:id/code", workspaceController.CreateWorkspaceCodeController, customMiddleware.Auth)
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
	userService := user.InitUserService(&app.Repository)
	jwtService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)

	customMiddleware := middlewares.InitMiddleware().WithUserService(userService).WithJwtTokenService(jwtService)

	e.GET("/:id", bookmarkController.FindBookmarkController, customMiddleware.Auth)
	e.GET("/:id/comments", bookmarkController.FindBookmarkCommentsController, customMiddleware.Auth)
	e.POST("", bookmarkController.CreateBookmarkController, customMiddleware.Auth)
	e.PUT("/:id", bookmarkController.UpdateBookmarkController, customMiddleware.Auth)
	e.DELETE("/:id", bookmarkController.DeleteBookmarkController, customMiddleware.Auth)
}

func (app *Application) InitRefreshTokenRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/refresh", V1))
	jwtTokenService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)
	refreshTokenService := refreshtoken.InitRefreshTokenService(&app.Repository, jwtTokenService)
	refreshTokenController := refreshtoken.InitRefreshTokenController().WithRefreshTokenService(refreshTokenService)

	e.POST("", refreshTokenController.RefreshAccessTokenController)
}

func (app *Application) InitCommentRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/comment", V1))
	categoryService := category.InitCategoryService().WithRepository(&app.Repository)
	invitecodeService := invitecode.InitInviteCodeService().WithRepository(&app.Repository)
	workspaceService := workspace.InitWorkspaceService(&app.Repository).WithInviteCodeService(&invitecodeService)
	invitecodeService = invitecodeService.WithWorkspaceService(&workspaceService)
	bookmarkService := bookmark.InitBookmarkService().WithRepository(&app.Repository).WithCategoryService(&categoryService)

	commentService := comment.InitCommentService().WithRepository(&app.Repository).WithBookmarkService(&bookmarkService)
	commentController := comment.InitCommentController().WithCommentService(&commentService)

	userService := user.InitUserService(&app.Repository)
	jwtService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)
	customMiddleware := middlewares.InitMiddleware().WithUserService(userService).WithJwtTokenService(jwtService)
	e.POST("", commentController.CreateCommentController, customMiddleware.Auth)
	e.PUT("/:id", commentController.UpdateCommentController, customMiddleware.Auth)
	e.DELETE("/:id", commentController.DeleteCommentController, customMiddleware.Auth)
}

func (app *Application) InitWorkspaceCodeRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/code", V1))
	categoryService := category.InitCategoryService().WithRepository(&app.Repository)
	workspaceService := workspace.InitWorkspaceService(&app.Repository)
	bookmarkSeervice := bookmark.InitBookmarkService().WithRepository(&app.Repository).WithCategoryService(&categoryService).WithWorkspaceService(workspaceService)

	gptRepository := gpt.InitGPTRepository().WithClient(app.GPTClient)
	gptService := gpt.InitGPTService().WithRepository(&gptRepository)
	deviceInfoService := deviceinfo.InitDeviceInfoService().WithRepository(&app.Repository)
	workspaceCodeService := workspacecode.InitWorkspaceCodeService().
		WithRepository(&app.Repository).WithBookmarkService(&bookmarkSeervice).WithCategoryService(&categoryService).WithWorkspaceService(workspaceService).WithDeviceInfoService(&deviceInfoService).WithGPTService(&gptService)
	workspaceCodeController := workspacecode.InitWorkspaceCodeController().WithWorkspaceCodeService(&workspaceCodeService)
	e.POST("/predict", workspaceCodeController.PredictCategoryController)
	e.GET("", workspaceCodeController.FindWorkspaceCodeController)
}
