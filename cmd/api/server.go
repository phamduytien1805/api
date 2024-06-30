package api

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/phamduytien1805/pkg/config"
)

type application struct {
	srv    *http.Server
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

func (app *application) serve() error {
	// Start the server

	app.srv = &http.Server{
		Addr: app.config.Web.Http.Server.Port,
		// Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		app.logger.Info("caught signal", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := app.gracefulStop(ctx)
		shutdownError <- err
	}()

	app.logger.Info("starting server", "addr", app.srv.Addr, "env", app.config.Env)

	err := app.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	app.logger.Info("stopped server", "addr", app.srv.Addr)

	return nil
}

func (app *application) gracefulStop(ctx context.Context) error {
	err := app.srv.Shutdown(ctx)
	if err != nil {
		return err
	}

	app.logger.Info("completing background tasks", "addr", app.srv.Addr)
	return nil
}
