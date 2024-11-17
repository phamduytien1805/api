package user

import (
	"context"
	"log/slog"

	"github.com/phamduytien1805/pkgmodule/config"
	"github.com/phamduytien1805/pkgmodule/hash_generator"
	"github.com/phamduytien1805/pkgmodule/id_generator"
	"github.com/phamduytien1805/pkgmodule/token"
	data_access "github.com/phamduytien1805/usermodule/data-access"
)

type UserServiceImpl struct {
	gateway    UserGateWay
	hashGen    *hash_generator.Argon2idHash
	logger     *slog.Logger
	tokenMaker token.Maker
	config     *config.Config
}

func NewUserServiceImpl(store data_access.Store, tokenMaker token.Maker, config *config.Config, logger *slog.Logger, hashGen *hash_generator.Argon2idHash) UserService {
	return &UserServiceImpl{
		gateway:    &UserGatewayImpl{store: store},
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

	createdUser, err := svc.gateway.CreateUserWithCredential(ctx, &User{
		ID:            ID,
		Username:      user.Username,
		Email:         user.Email,
		EmailVerified: false,
		State:         1,
	}, &UserCredential{
		HashedPassword: hashSaltCredential.Hash,
		Salt:           hashSaltCredential.Salt,
	}, func(createdUser *User) error {
		//TODO: add logic to send email verification
		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (svc *UserServiceImpl) AuthenticateUserBasic(ctx context.Context, userForm *BasicAuthForm) (*UserSession, error) {
	user, err := svc.gateway.GetUserByEmail(ctx, userForm.Email)
	if err != nil {
		svc.logger.Error("error getting user by email", err)
		return nil, ErrorUserInvalidAuthenticate
	}
	userCredential, err := svc.gateway.GetUserCredentialByUserId(ctx, user.ID)
	if err != nil {
		svc.logger.Error("error getting user credential", err)
		return nil, ErrorUserInvalidAuthenticate
	}

	if err = svc.hashGen.Compare(userCredential.HashedPassword, userCredential.Salt, userForm.Credential); err != nil {
		return nil, ErrorUserInvalidAuthenticate
	}

	accessToken, accessPayload, err := svc.tokenMaker.CreateToken(
		user.Username,
		svc.config.Token.AccessTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := svc.tokenMaker.CreateToken(
		user.Username,
		svc.config.Token.RefreshTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	return &UserSession{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  *user,
	}, nil
}
