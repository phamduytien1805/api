package user

import (
	"errors"
	"net/http"
	"phamduytien1805/pkg/http_helpers"
	"phamduytien1805/user/core"

	"github.com/go-chi/render"
)

func (s *HttpServer) registerUser(w http.ResponseWriter, r *http.Request) {
	createUserRequest := &core.CreateUserForm{}
	if err := render.DecodeJSON(r.Body, createUserRequest); err != nil {
		http_helpers.BadRequestResponse(w, r, err)
		return
	}

	if err := s.validator.Struct(createUserRequest); err != nil {
		http_helpers.FailedValidationResponse(w, r, err)
		return
	}

	user, err := s.userSvc.CreateUserWithCredential(r.Context(), createUserRequest)
	if err != nil {
		if errors.As(err, &core.ErrorUserResourceConflict) {
			http_helpers.EditConflictResponse(w, r, err)
			return
		}
		http_helpers.ServerErrorResponse(w, r, err)
		return
	}

	http_helpers.Ok(w, r, http.StatusCreated, user)
}

func (s *HttpServer) authenticateUserBasic(w http.ResponseWriter, r *http.Request) {
	basicAuthForm := &core.BasicAuthForm{}
	if err := render.DecodeJSON(r.Body, basicAuthForm); err != nil {
		http_helpers.BadRequestResponse(w, r, err)
		return
	}

	if err := s.validator.Struct(basicAuthForm); err != nil {
		http_helpers.FailedValidationResponse(w, r, err)
		return
	}

	userSession, err := s.userSvc.AuthenticateUserBasic(r.Context(), basicAuthForm)
	if err != nil {
		if errors.Is(err, core.ErrorUserInvalidAuthenticate) {
			http_helpers.InvalidAuthenticateResponse(w, r, err)
			return
		}
		http_helpers.ServerErrorResponse(w, r, err)
		return
	}

	http_helpers.Ok(w, r, http.StatusCreated, userSession)

}
