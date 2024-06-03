package app

import (
	"context"
	"fmt"
	"path/filepath"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitFirebaseApp(filename string) (*firebase.App, error) {
	ctx := context.Background()
	serviceAccountKeyFilePath, err := filepath.Abs(fmt.Sprintf("config/%s", filename))
	if err != nil {
		return nil, err
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
