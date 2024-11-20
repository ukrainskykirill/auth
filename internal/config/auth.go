package config

import (
	"os"
	"strconv"
)

const (
	refreshTokenTTLEnv = "JWT_REFRESH_TOKEN_TTL"
	accessTokenTTLEnv  = "JWT_ACCESS_TOKEN_TTL"
	tokenSecretEnv     = "JWT_TOKEN_SECRET"
)

type AuthConfig interface {
	RefreshTokenTTL() int
	AccessTokenTTL() int
	TokenSecret() string
}

type authConfig struct {
	refreshTokenTTL int
	accessTokenTTL  int
	tokenSecret     string
}

func NewAuthConfig() (AuthConfig, error) {
	tokenSecret, ok := os.LookupEnv(tokenSecretEnv)
	if !ok {
		return &authConfig{}, errVariableNotFound
	}

	refreshTokenTTLstr, ok := os.LookupEnv(refreshTokenTTLEnv)
	if !ok {
		return &authConfig{}, errVariableNotFound
	}
	refreshTokenTTL, err := strconv.Atoi(refreshTokenTTLstr)
	if err != nil {
		return &authConfig{}, errVariableParse
	}

	accessTokenTTLstr, ok := os.LookupEnv(accessTokenTTLEnv)
	if !ok {
		return &authConfig{}, errVariableNotFound
	}
	accessTokenTTL, err := strconv.Atoi(accessTokenTTLstr)
	if err != nil {
		return &authConfig{}, errVariableParse
	}

	return &authConfig{
		refreshTokenTTL: refreshTokenTTL,
		accessTokenTTL:  accessTokenTTL,
		tokenSecret:     tokenSecret,
	}, nil
}

func (c *authConfig) RefreshTokenTTL() int {
	return c.refreshTokenTTL
}

func (c *authConfig) AccessTokenTTL() int {
	return c.accessTokenTTL
}

func (c *authConfig) TokenSecret() string {
	return c.tokenSecret
}
