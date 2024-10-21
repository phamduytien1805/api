package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Route("/user", func(r chi.Router) {
		r.Post("/register", app.registerUser)
		r.Post("/auth", app.authenticateUserBasic)

	})

	r.Route("/chat", func(r chi.Router) {
		r.Get("/ws", app.wsHandler)
	})

	return r
}
