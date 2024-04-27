package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "devmark",
	Short: "Devmark is a bookmark service for developer.",
	Long: `Devmark is flexible bookmark service for those who want to organize
         their bookmark automatically.`,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
