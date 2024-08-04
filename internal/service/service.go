package service

import (
	"context"
	"github.com/ukrainskykirill/auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, user *model.UserIn) (int64, error)
	Get(ctx context.Context, userID int64) (*model.User, error)
	Update(ctx context.Context, user *model.UserInUpdate) error
	Delete(ctx context.Context, userID int64) error
}
