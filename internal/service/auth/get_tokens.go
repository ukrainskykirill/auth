package auth

import (
	"context"
	"fmt"

	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *authServ) GetTokens(ctx context.Context, oldRefreshToken string) (*model.TokensOut, error) {
	parsedRefreshClaims, err := s.parseRefreshToken(ctx, oldRefreshToken)
	if err != nil {
		return &model.TokensOut{}, err
	}
	fmt.Printf("parsed refresh token: %s", parsedRefreshClaims.Name)
	authInfo, err := s.userRepo.GetUserAuthInfo(ctx, parsedRefreshClaims.Name)
	if err != nil {
		return &model.TokensOut{}, err
	}
	fmt.Printf("parsed refresh token: %s", parsedRefreshClaims.Name)

	authInfoToClaims := model.AuthInfoToClaims{
		Name: parsedRefreshClaims.Name,
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
