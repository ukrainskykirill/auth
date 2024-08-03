package user

import (
	"github.com/ukrainskykirill/auth/internal/repository"
	"github.com/ukrainskykirill/auth/internal/service"
)

type userService struct {
	userRepo *repository.UserRepository
}

func NewService(repo *repository.UserRepository) service.UserService {
	return &userService{
		userRepo: repo,
	}
}
