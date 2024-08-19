package user

import (
	"context"
	"fmt"

	prError "github.com/ukrainskykirill/auth/internal/error"
	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *userServ) Get(ctx context.Context, userID int64) (*model.User, error) {
	isExist, err := s.repo.IsExistByID(ctx, userID)
	if err != nil {
		return &model.User{}, err
	}
	if !isExist {
		return &model.User{}, fmt.Errorf("service.Get - %w", prError.ErrUserNotFound)
	}

	user, err := s.repo.Get(ctx, userID)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}
