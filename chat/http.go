package chat

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"phamduytien1805/pkg/config"
	"phamduytien1805/pkg/http_helpers"
	"phamduytien1805/pkg/validator"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type HttpServer struct {
	httpServer *http.Server
	config     *config.Config
	logger     *slog.Logger
	validator  *validator.Validate
	httpPort   string
	router     *chi.Mux
}

func NewHttpServer(config *config.Config, logger *slog.Logger, validator *validator.Validate) *HttpServer {
	return &HttpServer{
		config:    config,
		logger:    logger,
		validator: validator,
		httpPort:  config.Web.Http.WS.Port,
	}
}

func (s *HttpServer) RegisterRoutes() {
	s.router = chi.NewRouter()
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Heartbeat("/ping"))

	s.router.NotFound(http_helpers.NotFoundResponse)
	s.router.MethodNotAllowed(http_helpers.MethodNotAllowedResponse)

}

func (s *HttpServer) Run() {
	go func() {
		addr := ":" + s.httpPort
		s.httpServer = &http.Server{
			Addr:    addr,
			Handler: s.router,
		}
		s.logger.Info("http server listening", slog.String("addr", addr))
		err := s.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.logger.Error(err.Error())
			os.Exit(1)
		}
	}()
}

func (r *HttpServer) GracefulStop(ctx context.Context) error {

	err := r.httpServer.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
