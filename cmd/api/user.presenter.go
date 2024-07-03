package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/phamduytien1805/internal/user"
)

type CreateUserRequest struct {
	Username   string `json:"username" validate:"required,min=5,max=32"`
	Email      string `json:"email" validate:"required,email"`
	Credential string `json:"credential" validate:"required,min=9"`
}

type CreateUserResponse struct {
	*user.User
}

func (rd *CreateUserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}
