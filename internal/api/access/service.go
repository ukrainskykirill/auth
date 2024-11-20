package access

import (
	"github.com/ukrainskykirill/auth/internal/service"
	gaccess "github.com/ukrainskykirill/auth/pkg/access_v1"
)

type Implementation struct {
	gaccess.UnimplementedAccessV1Server
	authService service.AuthService
}

func NewAccessImplementation(authService service.AuthService) *Implementation {
	return &Implementation{authService: authService}
}
