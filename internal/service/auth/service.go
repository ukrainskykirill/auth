package auth

import (
	"github.com/ukrainskykirill/auth/internal/cache"
	"github.com/ukrainskykirill/auth/internal/repository"
	"github.com/ukrainskykirill/auth/internal/service"
)

type authServ struct {
	userRepo        repository.UserRepository
	accessRepo      repository.AccessRepository
	cache           cache.AuthCache
	tokenSecret     string
	accessTokenTTL  int
	refreshTokenTTL int
}

func NewAuthServ(
	userRepo repository.UserRepository,
	accessRepo repository.AccessRepository,
	cache cache.AuthCache,
	tokenSecret string,
	accessTokenTTL int,
	refreshTokenTTL int,
) service.AuthService {
	return &authServ{
		userRepo:        userRepo,
		accessRepo:      accessRepo,
		cache:           cache,
		tokenSecret:     tokenSecret,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}
