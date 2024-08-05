package user

import (
	"context"
	"errors"
	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *userServ) Get(ctx context.Context, userID int64) (*model.User, error) {
	isExist, err := s.repo.IsExistByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return &model.User{}, errors.New("user not found")
	}

	user, err := s.repo.Get(ctx, userID)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}
