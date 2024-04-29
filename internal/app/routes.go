package app

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/auth"
	"github.com/MJU-Capstone-6/devmark-backend/internal/middlewares"
	"github.com/MJU-Capstone-6/devmark-backend/internal/user"
	"github.com/labstack/echo/v4/middleware"
)

func (app *Application) InitRoutes() {
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))
	app.Handler.Use(middlewares.ParseHeader)
	app.InitUserRoutes()
	app.InitAuthRoutes()
}

func (app *Application) InitUserRoutes() {
	e := app.Handler.Group("/user")
	userController := user.InitController(app.DB)
	e.POST("", userController.CreateUser)
}

func (app *Application) InitAuthRoutes() {
	e := app.Handler.Group("/auth")
	authController := auth.InitAuthController(app.DB, app.Config.Kakao)
	e.GET("/:provider", authController.GetKakaoUserInfo)
}
