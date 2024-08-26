package user

import (
	"context"

	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *userServ) Get(ctx context.Context, userID int64) (*model.User, error) {
	user, err := s.cache.Get(ctx, userID)
	if err != nil {
		return &model.User{}, err
	}

	if user != nil && *user == (model.User{}) {
		user, err := s.repo.Get(ctx, userID)
		if err != nil {
			return &model.User{}, err
		}

		err = s.cache.Create(ctx, user)
		if err != nil {
			return &model.User{}, err
		}
	}
	return user, nil
}
