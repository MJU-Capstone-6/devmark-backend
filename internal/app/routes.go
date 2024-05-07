package app

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/auth"
	"github.com/MJU-Capstone-6/devmark-backend/internal/jwtToken"
	"github.com/MJU-Capstone-6/devmark-backend/internal/middlewares"
	"github.com/MJU-Capstone-6/devmark-backend/internal/user"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (app *Application) InitRoutes() {
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))
	app.Handler.Use(middlewares.ParseHeader)
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
	authController := auth.InitAuthController().
		WithUserService(userService).
		WithAuthService(app.DB).
		WithKakaoInfo(app.Config.Kakao).
		WithJWTService(jwtService)
	e.GET("/:provider", authController.GetKakaoUserInfo)
}
