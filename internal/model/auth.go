package model

import "github.com/golang-jwt/jwt/v5"

type LoginIn struct {
	Name     string
	Password string
}

type TokensOut struct {
	AccessToken  string
	RefreshToken string
}
type AuthInfoToClaims struct {
	Name string
	Role string
}

type UserAuthInfo struct {
	Password string
	Role     string
}

type UserAccessClaims struct {
	jwt.RegisteredClaims
	Name string
	Role string
}

type UserRefreshClaims struct {
	jwt.RegisteredClaims
	Name string
}
