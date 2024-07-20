package user

import "context"

type UserService interface {
	CreateUserWithCredential(ctx context.Context, user *CreateUserForm) (*User, error)
}
