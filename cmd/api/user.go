package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/phamduytien1805/internal/user"
	user_pkg "github.com/phamduytien1805/internal/user"
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
		if errors.As(err, &user_pkg.ErrorUserResourceConflict) {
			app.editConflictResponse(w, r, err)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	app.Ok(w, r, http.StatusCreated, user)
}

func (app *application) authenticateUserBasic(w http.ResponseWriter, r *http.Request) {
	basicAuthForm := &user.BasicAuthForm{}
	if err := render.DecodeJSON(r.Body, basicAuthForm); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.validator.Struct(basicAuthForm); err != nil {
		app.failedValidationResponse(w, r, err)
		return
	}
	app.Ok(w, r, http.StatusCreated, nil)

}
