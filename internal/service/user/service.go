package user

import (
	"github.com/ukrainskykirill/auth/internal/cache"
	"github.com/ukrainskykirill/auth/internal/repository"
	"github.com/ukrainskykirill/auth/internal/service"
)

type userServ struct {
	repo  repository.UserRepository
	cache cache.UserCache
}

func NewServ(userRepo repository.UserRepository, cache cache.UserCache) service.UserService {
	return &userServ{
		repo: userRepo, cache: cache,
	}
}
