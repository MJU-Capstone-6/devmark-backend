package cmd

import (
	"log"

	"github.com/MJU-Capstone-6/devmark-backend/internal/app"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Puddlee Backend Server",
	Long: `Run Puddlee Backend Server on local.
	configuration file lies on the config directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Run() error {
	application, err := app.InitApplication()
	if err != nil {
		return err
	}
	application.Run()
	return nil
}
