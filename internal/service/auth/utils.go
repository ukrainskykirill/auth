package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"

	prError "github.com/ukrainskykirill/auth/internal/error"
	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *authServ) createAccessTokenClaims(ctx context.Context, authInfo *model.AuthInfoToClaims) *model.UserAccessClaims {
	return &model.UserAccessClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Second * time.Duration(s.accessTokenTTL)),
			},
		},
		Name: authInfo.Name,
		Role: authInfo.Role,
	}

}

func (s *authServ) createRefreshTokenClaims(ctx context.Context, authInfo *model.AuthInfoToClaims) *model.UserRefreshClaims {
	return &model.UserRefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Second * time.Duration(s.refreshTokenTTL)),
			},
		},
		Name: authInfo.Name,
	}

}

func (s *authServ) createTokensPairWithClaims(ctx context.Context, accessClaims *model.UserAccessClaims, refreshClaims *model.UserRefreshClaims) (*model.TokensOut, error) {
	accessTokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := accessTokenWithClaims.SignedString([]byte(s.tokenSecret))
	if err != nil {
		return &model.TokensOut{}, err
	}

	refreshTokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshTokenWithClaims.SignedString([]byte(s.tokenSecret))
	if err != nil {
		return &model.TokensOut{}, err
	}

	return &model.TokensOut{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func (s *authServ) ParseAccessToken(accessToken string) (*model.UserAccessClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(
		accessToken, &model.UserAccessClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.tokenSecret), nil
		})
	if err != nil {
		return &model.UserAccessClaims{}, err
	}

	if !parsedAccessToken.Valid {
		return &model.UserAccessClaims{}, prError.ErrInvalidToken
	}

	return parsedAccessToken.Claims.(*model.UserAccessClaims), nil
}

func (s *authServ) parseRefreshToken(ctx context.Context, refreshToken string) (*model.UserRefreshClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(
		refreshToken, &model.UserRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.tokenSecret), nil
		})
	if err != nil {
		return &model.UserRefreshClaims{}, err
	}

	if !parsedRefreshToken.Valid {
		return &model.UserRefreshClaims{}, prError.ErrInvalidToken
	}

	return parsedRefreshToken.Claims.(*model.UserRefreshClaims), nil
}
