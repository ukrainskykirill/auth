package cache

import (
	"context"

	"github.com/ukrainskykirill/auth/internal/model"
)

type UserCache interface {
	Create(ctx context.Context, userIn *model.User) error
	Get(ctx context.Context, id int64) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}
