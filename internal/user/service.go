package user

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	data_access "github.com/phamduytien1805/internal/data-access"
	"github.com/phamduytien1805/pkg/common"
	"github.com/phamduytien1805/pkg/hash_generator"
	"github.com/phamduytien1805/pkg/id_generator"
)

type UserServiceImpl struct {
	store   data_access.Store
	hashGen *hash_generator.Argon2idHash
}

func NewUserServiceImpl(store data_access.Store, hashGen *hash_generator.Argon2idHash) UserService {
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

	hashSaltCredential, err := svc.hashGen.GenerateHash([]byte(user.Credential), nil)

	arg := data_access.CreateUserWithCredentialTxParams{
		CreateUserParams: data_access.CreateUserParams{
			ID:            ID,
			Username:      user.Username,
			Email:         user.Email,
			EmailVerified: false,
			State:         1,
		},
		HashedCredential: hex.EncodeToString(hashSaltCredential.Hash),
		Salt:             hex.EncodeToString(hashSaltCredential.Salt),
		AfterCreate: func(createdUser data_access.User) error {
			//TODO: add logic to send email verification
			return nil
		},
	}
	txResult, err := svc.store.CreateUserWithCredentialTx(ctx, arg)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == common.UNIQUE_CONSTRAINT_VIOLATION {
				return nil, ErrorUserResourceConflict
			}

		}
		return nil, err
	}

	return mapToUser(txResult.User), nil
}

func (svc *UserServiceImpl) AuthenticateUserBasic(ctx context.Context, user *BasicAuthForm) (*User, error) {
	// userData, err := svc.store.GetUserByEmail(ctx, user.Email)
	// if err != nil {
	// 	return nil, err
	// }

	// hashSaltCredential := hash_generator.HashSalt{
	// 	Hash: []byte(userData.HashedCredential),
	// 	Salt: []byte(userData.Salt),
	// }

	// if !svc.hashGen.CompareHash([]byte(user.Credential), hashSaltCredential) {
	// 	return nil, ErrorUserInvalidCredential
	// }

	return nil, nil
}
