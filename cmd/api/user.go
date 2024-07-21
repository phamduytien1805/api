package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/phamduytien1805/internal/user"
)

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	createUserRequest := &user.CreateUserForm{}
	if err := render.DecodeJSON(r.Body, createUserRequest); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.validator.Struct(createUserRequest); err != nil {
		app.failedValidationResponse(w, r, err)
		return
	}

	user, err := app.userSvc.CreateUserWithCredential(r.Context(), createUserRequest)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.Ok(w, r, http.StatusCreated, user)
}
