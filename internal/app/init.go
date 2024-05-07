package app

import (
	"context"
	"crypto/ed25519"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/MJU-Capstone-6/devmark-backend/internal/config"
	"github.com/MJU-Capstone-6/devmark-backend/internal/db"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/labstack/echo/v4"
)

var (
	app  *Application
	once sync.Once
)

func InitApplication() (*Application, error) {
	if app == nil {
		once.Do(func() {
			err := setApplication()
			if err != nil {
				log.Fatal("something went wrong while configure application")
			}
		})
	} else {
		return nil, errors.New("Application already configured")
	}
	return app, nil
}

func (app *Application) Run() error {
	if err := app.Handler.Start(fmt.Sprintf(":%s", app.Config.App.Port)); err != nil {
		return err
	}
	defer app.DB.Close(context.Background())
	return nil
}

func setApplication() error {
	ctx := context.Background()
	applicationConfig, err := config.InitConfig()
	if err != nil {
		return err
	}
	db, err := db.InitDB(ctx, applicationConfig.DB)
	if err != nil {
		return err
	}
	handler := echo.New()
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return err
	}
	app = &Application{
		DB:         db,
		Config:     applicationConfig,
		Repository: *repository.New(db),
		PubKey:     publicKey,
		PrivateKey: privateKey,
		Handler:    handler,
	}

	app.InitRoutes()

	return nil
}
