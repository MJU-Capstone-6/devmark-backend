package app

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/user"
)

func (app *Application) InitRoutes() {
	app.InitUserRoutes()
}

func (app *Application) InitUserRoutes() {
	e := app.Handler.Group("/user")
	userController := user.InitController(app.DB)
	e.POST("", userController.CreateUser)
}
