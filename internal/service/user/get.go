package user

import (
	"context"
	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *userService) Get(ctx context.Context, userID int64) (*model.User, error) {
	return &model.User{}, nil
}
