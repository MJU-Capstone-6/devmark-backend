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

func (app *Application) InitSwaggerRoutes() {
}

func (app *Application) InitUserRoutes() {
}

func (app *Application) InitAuthRoutes() {
	e := app.Handler.Group("/auth")
	authController := auth.InitAuthController(app.DB, app.Config.Kakao)
	e.GET("/:provider", authController.GetKakaoUserInfo)
}
