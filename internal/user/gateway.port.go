package user

import (
	"context"

	"github.com/google/uuid"
)

type UserGateWay interface {
	CreateUserWithCredential(ctx context.Context, userParams *User, userCredential *UserCredential, afterCreateFn func(*User) error) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserCredentialByUserId(ctx context.Context, userID uuid.UUID) (*UserCredential, error)
}
