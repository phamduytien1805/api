package user

import (
	data_access "github.com/phamduytien1805/internal/data-access"
)

func mapToUser(u data_access.User) *User {
	return &User{
		ID:            u.ID,
		Username:      u.Username,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		State:         u.State,
	}
}

func mapToUserCredential(uc data_access.UserCredential) *UserCredential {
	return &UserCredential{
		HashedPassword: uc.Credential,
		Salt:           uc.Salt,
	}
}
