package user

import (
	"github.com/google/uuid"
	data_access "github.com/phamduytien1805/internal/data-access"
)

type CreateUserForm struct {
	Username   string `json:"username" validate:"required,min=5,max=32"`
	Email      string `json:"email" validate:"required,email"`
	Credential string `json:"credential" validate:"required,min=9"`
}

type User struct {
	ID            uuid.UUID
	Username      string
	Email         string
	EmailVerified bool
}

func mapToUser(u data_access.User) *User {
	return &User{
		ID:            u.ID,
		Username:      u.Username,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
	}
}
