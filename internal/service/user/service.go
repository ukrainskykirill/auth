package user

import (
	"github.com/ukrainskykirill/auth/internal/repository"
	"github.com/ukrainskykirill/auth/internal/service"
)

type userServ struct {
	repo repository.UserRepository
}

func NewServ(userRepo repository.UserRepository) service.UserService {
	return &userServ{
		repo: userRepo,
	}
}
