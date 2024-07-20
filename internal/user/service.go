package user

import (
	"context"

	data_access "github.com/phamduytien1805/internal/data-access"
	"github.com/phamduytien1805/pkg/id_generator"
)

type UserServiceImpl struct {
	store data_access.Store
}

func NewUserServiceImpl(store data_access.Store) *UserServiceImpl {
	return &UserServiceImpl{
		store: store,
	}
}

func (svc *UserServiceImpl) CreateUserWithCredential(ctx context.Context, user *CreateUserForm) (*User, error) {
	ID, err := id_generator.NewUUID()
	if err != nil {
		return nil, err
	}
	arg := data_access.CreateUserTxParams{
		CreateUserParams: data_access.CreateUserParams{
			ID:            ID,
			Username:      user.Username,
			Email:         user.Email,
			EmailVerified: false,
			State:         1,
		},
		AfterCreate: func(user data_access.User) error {
			_, err := svc.store.CreateUserCredential(ctx, data_access.CreateUserCredentialParams{
				UserID:     user.ID,
				Credential: "",
				Salt:       "user.Salt",
			})
			return err
		},
	}
	txResult, err := svc.store.CreateUserTx(ctx, arg)
	return mapToUser(txResult.User), err
}
