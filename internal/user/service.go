package user

import (
	"context"
	"encoding/hex"

	data_access "github.com/phamduytien1805/internal/data-access"
	"github.com/phamduytien1805/pkg/hash_generator"
	"github.com/phamduytien1805/pkg/id_generator"
)

type UserServiceImpl struct {
	store   data_access.Store
	hashGen *hash_generator.Argon2idHash
}

func NewUserServiceImpl(store data_access.Store, hashGen *hash_generator.Argon2idHash) *UserServiceImpl {
	return &UserServiceImpl{
		store:   store,
		hashGen: hashGen,
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
		AfterCreate: func(createdUser data_access.User) error {
			hashSalt, err := svc.hashGen.GenerateHash([]byte(user.Credential), nil)
			if err != nil {
				return err
			}
			_, err = svc.store.CreateUserCredential(ctx, data_access.CreateUserCredentialParams{
				UserID:     createdUser.ID,
				Credential: hex.EncodeToString(hashSalt.Hash),
				Salt:       hex.EncodeToString(hashSalt.Salt),
			})
			return err
		},
	}
	txResult, err := svc.store.CreateUserTx(ctx, arg)
	return mapToUser(txResult.User), err
}
