// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package data_access

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserCredential(ctx context.Context, arg CreateUserCredentialParams) (UserCredential, error)
	CreateUserSocial(ctx context.Context, arg CreateUserSocialParams) (UserSocialToken, error)
}

var _ Querier = (*Queries)(nil)
