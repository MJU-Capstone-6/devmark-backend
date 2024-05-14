package db

import (
	"context"
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

	_, err = client.SchemaApply(context.Background(), &atlasexec.SchemaApplyParams{
		URL: dbURL,
	})
	if err != nil {
		return err
	}
	return nil
}
