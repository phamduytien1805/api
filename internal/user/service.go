package user

import (
	"context"
	"encoding/hex"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
	data_access "github.com/phamduytien1805/internal/data-access"
	"github.com/phamduytien1805/pkg/common"
	"github.com/phamduytien1805/pkg/config"
	"github.com/phamduytien1805/pkg/hash_generator"
	"github.com/phamduytien1805/pkg/id_generator"
	"github.com/phamduytien1805/pkg/token"
)

type UserServiceImpl struct {
	store      data_access.Store
	hashGen    *hash_generator.Argon2idHash
	logger     *slog.Logger
	tokenMaker token.Maker
	config     *config.Config
}

func NewUserServiceImpl(store data_access.Store, tokenMaker token.Maker, config *config.Config, logger *slog.Logger, hashGen *hash_generator.Argon2idHash) UserService {
	return &UserServiceImpl{
		store:      store,
		logger:     logger,
		hashGen:    hashGen,
		tokenMaker: tokenMaker,
		config:     config,
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

func (svc *UserServiceImpl) AuthenticateUserBasic(ctx context.Context, userForm *BasicAuthForm) (*User, error) {
	user, err := svc.store.GetUserByEmail(ctx, userForm.Email)
	if err != nil {
		svc.logger.Error("error getting user by email", err)
		return nil, ErrorUserInvalidAuthenticate
	}
	userCredential, err := svc.store.GetUserCredentialByUserId(ctx, user.ID)
	if err != nil {
		svc.logger.Error("error getting user credential", err)
		return nil, ErrorUserInvalidAuthenticate
	}

	if err = svc.hashGen.Compare([]byte(userCredential.Credential), []byte(userCredential.Salt), []byte(userForm.Credential)); err != nil {
		return nil, ErrorUserInvalidAuthenticate
	}
	accessToken, accessPayload, err := svc.tokenMaker.CreateToken(
		user.Username,
		svc.config.Token.AccessTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
