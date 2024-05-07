package app

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/auth"
	"github.com/MJU-Capstone-6/devmark-backend/internal/jwtToken"
	"github.com/MJU-Capstone-6/devmark-backend/internal/middlewares"
	refreshtoken "github.com/MJU-Capstone-6/devmark-backend/internal/refreshToken"
	"github.com/MJU-Capstone-6/devmark-backend/internal/user"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (app *Application) InitRoutes() {
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))
	app.Handler.File("/docs", "swagger.json")
	app.Handler.GET("/swagger/*", echoSwagger.WrapHandler)
	app.InitUserRoutes()
	app.InitAuthRoutes()
}

func (app *Application) InitUserRoutes() {
}

func (app *Application) InitAuthRoutes() {
	e := app.Handler.Group("/auth")
	userService := user.InitUserService(&app.Repository)
	jwtService := jwtToken.InitJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)
	refreshTokenService := refreshtoken.InitRefreshTokenService(&app.Repository, jwtService)
	authService := auth.InitAuthService(&app.Repository, userService, jwtService, refreshTokenService)
	authController := auth.InitAuthController().
		WithKakaoInfo(app.Config.Kakao).
		WithAuthService(authService)

	e.POST("/:provider", authController.GetKakaoUserInfo, middlewares.ParseHeader)
}
