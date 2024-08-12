package repository

import (
	"context"
	"github.com/ukrainskykirill/auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.UserIn) (int64, error)
	Get(ctx context.Context, userID int64) (*model.User, error)
	Update(ctx context.Context, user *model.UserInUpdate) error
	Delete(ctx context.Context, userID int64) error
	IsExistByID(ctx context.Context, userID int64) (bool, error)
	IsExistByName(ctx context.Context, name string) (bool, error)
}
