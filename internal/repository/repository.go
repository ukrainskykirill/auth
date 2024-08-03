package repository

import (
	"context"
	"github.com/ukrainskykirill/auth/internal/model"
	modelRepo "github.com/ukrainskykirill/auth/internal/repository/user/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *modelRepo.RepoUserIn) (int64, error)
	Get(ctx context.Context, userID int64) (*model.User, error)
	Update(ctx context.Context, user *modelRepo.RepoUserInUpdate) error
	Delete(ctx context.Context, userID int64) error
}
