package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/phamduytien1805/chatmodule/internal/chat"
	"github.com/phamduytien1805/pkgmodule/config"
)

type application struct {
	srv    *http.Server
	logger *slog.Logger
	config *config.Config
	hub    *chat.Hub
}

func initializeApplication() (*application, error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	hub := chat.NewHub()

	app := &application{
		logger: logger,
		config: configConfig,
		hub:    hub,
	}
	return app, nil
}

func (app *application) serve() error {
	// Start the server

	app.srv = &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.Web.Http.Server.Port),
		Handler:      app.routes(),
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

	go app.hub.Run()

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
	return nil
}
