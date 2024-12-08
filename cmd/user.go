package cmd

import (
	"log/slog"
	"os"
	"phamduytien1805/pkg/common"
	"phamduytien1805/pkg/config"
	"phamduytien1805/pkg/db"
	"phamduytien1805/pkg/hash_generator"
	"phamduytien1805/pkg/token"
	"phamduytien1805/pkg/validator"
	"phamduytien1805/user"
	"phamduytien1805/user/core"
	"phamduytien1805/user/data_access"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "user",
	Short: "api server",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := initializeUserServer()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		app.Serve()
	},
}

func initializeUserServer() (*common.Server, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	validator := validator.New()

	db, err := db.NewPostgresql(configConfig)

	if err != nil {
		return nil, err
	}
	store := data_access.NewStore(db)
	hashGen := hash_generator.NewArgon2idHash(configConfig)
	tokenMaker, err := token.NewJWTMaker(configConfig.Token.SecretKey)
	if err != nil {
		return nil, err
	}

	userSvc := core.NewUserServiceImpl(store, tokenMaker, configConfig, logger, hashGen)

	httpServer := user.NewHttpServer(configConfig, logger, validator, userSvc)
	router := user.NewRouter(httpServer)

	return common.NewServer(router, nil), nil
}

func init() {
	RootCmd.AddCommand(apiCmd)
}
