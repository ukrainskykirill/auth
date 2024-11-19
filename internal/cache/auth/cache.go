package auth

import (
	platformCache "github.com/ukrainskykirill/platform_common/pkg/cache"

	"github.com/ukrainskykirill/auth/internal/cache"
)

type AuthImplementation struct {
	cl platformCache.Client
}

func NewCache(cl platformCache.Client) cache.AuthCache {
	return &AuthImplementation{
		cl: cl,
	}
}
