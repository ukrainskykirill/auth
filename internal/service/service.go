package service

import (
	"context"

	"github.com/ukrainskykirill/auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, user *model.UserIn) (int64, error)
	Get(ctx context.Context, userID int64) (*model.User, error)
	Update(ctx context.Context, user *model.UserInUpdate) error
	Delete(ctx context.Context, userID int64) error
}

type AuthService interface {
	Login(ctx context.Context, loginIn *model.LoginIn) (*model.TokensOut, error)
	GetTokens(ctx context.Context, oldRefreshToken string) (*model.TokensOut, error)
	Check(ctx context.Context, role, endpoint string) error
	ParseAccessToken(accessToken string) (*model.UserAccessClaims, error)
}
