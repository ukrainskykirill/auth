package auth

import (
	"context"

	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *authServ) Login(ctx context.Context, loginIn *model.LoginIn) (*model.TokensOut, error) {
	return &model.TokensOut{}, nil
}

func (s *authServ) GetTokens(ctx context.Context, oldRefreshToken string) (*model.TokensOut, error) {
	return &model.TokensOut{}, nil
}

func (s *authServ) Check(ctx context.Context, endpoint string) error {
	return nil
}
