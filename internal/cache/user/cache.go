package user

 import (
 	platformCache "github.com/ukrainskykirill/platform_common/pkg/cache"

 	"github.com/ukrainskykirill/auth/internal/cache"
 )

 type Implementation struct {
 	cl platformCache.Client
 }

 func NewCache(cl platformCache.Client) cache.UserCache {
 	return &Implementation{
 		cl: cl,
 	}
 }