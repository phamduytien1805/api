package api

import (
	"net/http"

	"github.com/go-chi/render"
)

func (app *application) Ok(w http.ResponseWriter, r *http.Request, payload render.Renderer) {
	render.Render(w, r, payload)
}
