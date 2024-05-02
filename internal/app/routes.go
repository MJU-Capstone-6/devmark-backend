package app

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/auth"
	"github.com/MJU-Capstone-6/devmark-backend/internal/middlewares"
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
	authController := auth.InitAuthController().
		WithUserService(app.DB).
		WithAuthService(app.DB).
		WithKakaoInfo(app.Config.Kakao).
		WithJWTService(app.PubKey, app.PrivateKey, app.Config.App.FooterKey)
	e.GET("/:provider", authController.GetKakaoUserInfo)
}
