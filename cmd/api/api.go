package api

import (
	"log/slog"
	"os"

	"github.com/phamduytien1805/cmd"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "api server",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := initializeApplication()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		app.serve()

	},
}

func init() {
	cmd.RootCmd.AddCommand(apiCmd)
}
