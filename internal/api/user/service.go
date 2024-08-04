package user

import (
	"github.com/ukrainskykirill/auth/internal/service"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
)

type Implementation struct {
	guser.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{userService: userService}
}
