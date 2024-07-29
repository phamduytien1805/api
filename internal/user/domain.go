package user

import (
	"github.com/google/uuid"
	data_access "github.com/phamduytien1805/internal/data-access"
)

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
