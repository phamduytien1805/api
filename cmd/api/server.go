package api

import (
	"log/slog"
	"os"

	"github.com/phamduytien1805/pkg/config"
)

type application struct {
	config *config.Config
	logger *slog.Logger
}

func initializeApplication() (*application, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		config: configConfig,
		logger: logger,
	}
	return app, nil
}

func (app *application) serve() {
	// Start the server

	// srv := &http.Server{
	// 	Addr: app.config.Web.Http.Server.Port,
	// 	// Handler:      app.routes(),
	// 	IdleTimeout:  time.Minute,
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }
}
