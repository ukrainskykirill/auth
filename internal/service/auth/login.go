package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *authServ) Login(ctx context.Context, loginIn *model.LoginIn) (*model.TokensOut, error) {
	authInfo, err := s.userRepo.GetUserAuthInfo(ctx, loginIn.Name)
	if err != nil {
		return &model.TokensOut{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(authInfo.Password), []byte(loginIn.Password))
	if err != nil {
		return &model.TokensOut{}, err
	}

	authInfoToClaims := model.AuthInfoToClaims{
		Name: loginIn.Name,
		Role: authInfo.Role,
	}

	accessClaims := s.createAccessTokenClaims(ctx, &authInfoToClaims)
	refreshClaims := s.createRefreshTokenClaims(ctx, &authInfoToClaims)
	tokens, err := s.createTokensPairWithClaims(ctx, accessClaims, refreshClaims)
	if err != nil {
		return &model.TokensOut{}, err
	}

	return &model.TokensOut{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
