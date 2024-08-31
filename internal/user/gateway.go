package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	data_access "github.com/phamduytien1805/internal/data-access"
	"github.com/phamduytien1805/pkg/common"
)

type UserGatewayImpl struct {
	store data_access.Store
}

func NewUserGatewayImpl(store data_access.Store) UserGateWay {
	return &UserGatewayImpl{
		store: store,
	}
}

func (gw *UserGatewayImpl) CreateUserWithCredential(ctx context.Context, userParams *User, userCredential *UserCredential, afterCreateFn func(*User) error) (*User, error) {
	arg := data_access.CreateUserWithCredentialTxParams{
		CreateUserParams: data_access.CreateUserParams{
			ID:            userParams.ID,
			Username:      userParams.Username,
			Email:         userParams.Email,
			EmailVerified: userParams.EmailVerified,
			State:         userParams.State,
		},
		HashedCredential: userCredential.HashedPassword,
		Salt:             userCredential.Salt,
		AfterCreate: func(u data_access.User) error {
			return afterCreateFn(mapToUser(u))
		},
	}
	txResult, err := gw.store.CreateUserWithCredentialTx(ctx, arg)
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

func (gw *UserGatewayImpl) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u, err := gw.store.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return mapToUser(u), nil
}

func (gw *UserGatewayImpl) GetUserCredentialByUserId(ctx context.Context, userID uuid.UUID) (*UserCredential, error) {
	uc, err := gw.store.GetUserCredentialByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}
	return mapToUserCredential(uc), nil
}
