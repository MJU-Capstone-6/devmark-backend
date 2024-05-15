package db

import (
	"context"
	"log"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"
)

func Migration(dbURL string) error {
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS("./migration"),
		),
	)
	if err != nil {
		return err
	}
	defer workdir.Close()

	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		return err
	}
	_, err = client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL: dbURL,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
