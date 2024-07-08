package api

import (
	"net/http"

	"github.com/go-chi/render"
)

// type envelope map[string]any

// func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
// 	js, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}

// 	js = append(js, '\n')

// 	for k, v := range headers {
// 		w.Header()[k] = v
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	w.Write(js)
// 	return nil
// }

func (app *application) Ok(w http.ResponseWriter, r *http.Request, payload render.Renderer) {
	render.Render(w, r, payload)
}
