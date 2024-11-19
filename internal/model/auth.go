package model

type LoginIn struct {
	name     string
	password string
}

type TokensOut struct {
	accessToken  string
	refreshToken string
}
