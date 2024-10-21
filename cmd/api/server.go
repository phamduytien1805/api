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

	data_access "github.com/phamduytien1805/internal/data-access"
	"github.com/phamduytien1805/internal/user"
	"github.com/phamduytien1805/pkg/config"
	"github.com/phamduytien1805/pkg/db"
	"github.com/phamduytien1805/pkg/hash_generator"
	"github.com/phamduytien1805/pkg/token"
	v "github.com/phamduytien1805/pkg/validator"
)

type application struct {
	srv       *http.Server
	config    *config.Config
	logger    *slog.Logger
	validator *v.Validate
	ws        WSConn
	userSvc   user.UserService
}

func initializeApplication() (*application, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	validator := v.New()

	ws := WebSocketBuilder(configConfig)

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

	userSvc := user.NewUserServiceImpl(store, tokenMaker, configConfig, logger, hashGen)

	app := &application{
		ws:        ws,
		config:    configConfig,
		logger:    logger,
		validator: validator,
		userSvc:   userSvc,
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
