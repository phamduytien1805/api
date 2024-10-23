package api

import (
	"log/slog"
	"os"

	"github.com/phamduytien1805/usermodule/cmd"
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
		err = app.serve()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

	},
}

func init() {
	cmd.RootCmd.AddCommand(apiCmd)
}
