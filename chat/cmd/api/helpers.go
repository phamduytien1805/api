package api

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) Ok(w http.ResponseWriter, r *http.Request, status int, payload any) {
	err := app.writeJSON(w, status, envelope{"data": payload, "success": true}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
