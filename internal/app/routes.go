package app

import (
	"fmt"

	"github.com/MJU-Capstone-6/devmark-backend/internal/auth"
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
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Handler.File("/docs", "swagger.json")
	app.Handler.GET("/swagger/*", echoSwagger.WrapHandler)
	app.InitUserRoutes()
	app.InitAuthRoutes()
	app.InitWorkspaceRoutes()
	app.InitInviteCodeRoutes()
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

	e.POST("/:provider", authController.GetKakaoUserInfo, middlewares.ParseHeader)
}

func (app *Application) InitWorkspaceRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/workspace", V1))
	workspaceService := workspace.InitWorkspaceService(&app.Repository)
	inviteCodeService := invitecode.InitInviteCodeService().WithRepository(&app.Repository).WithWorkspaceService(workspaceService)
	workspaceService.WithInviteCodeService(&inviteCodeService)
	workspaceController := workspace.InitWorkspaceController().WithWorkspaceService(workspaceService)

	e.GET("/:id", workspaceController.ViewWorkspaceController)
	e.PUT("/:id", workspaceController.UpdateWorkspaceController)
	e.POST("", workspaceController.CreateWorkspaceController)
	e.DELETE("/:id", workspaceController.DeleteWorkspaceController)
}

func (app *Application) InitInviteCodeRoutes() {
	e := app.Handler.Group(fmt.Sprintf("%s/invitecode", V1))
	workspaceService := workspace.InitWorkspaceService(&app.Repository)
	inviteCodeService := invitecode.InitInviteCodeService().
		WithRepository(&app.Repository).
		WithWorkspaceService(workspaceService)
	inviteCodeController := invitecode.InitInviteCodeController().WithInviteCodeService(&inviteCodeService)

	e.POST("", inviteCodeController.GenerateInviteCodeController)
}
