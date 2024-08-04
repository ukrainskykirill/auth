package user

import (
	"context"
	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *userServ) Get(ctx context.Context, userID int64) (*model.User, error) {
	user, err := s.repo.Get(ctx, userID)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}
