package auth

import (
	"github.com/ukrainskykirill/auth/internal/service"
	gauth "github.com/ukrainskykirill/auth/pkg/auth_v1"
)

type Implementation struct {
	gauth.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewAuthImplementation(authService service.AuthService) *Implementation {
	return &Implementation{authService: authService}
}
