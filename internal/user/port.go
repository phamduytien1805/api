package user

import "context"

type UserService interface {
	CreateUserWithCredential(ctx context.Context, user *CreateUserForm) (*User, error)
	AuthenticateUserBasic(ctx context.Context, user *BasicAuthForm) (*User, error)
}
