package user

import (
	"context"
	"fmt"

	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *userServ) Update(ctx context.Context, user *model.UserInUpdate) error {
	if len(user.Email) != 0 {
		err := validateEmail(user.Email)
		if err != nil {
			return fmt.Errorf("service.Update - %w", err)
		}
	}

	err := s.repo.Update(ctx, user)
	if err != nil {
		return err
	}

	err = s.cache.Delete(ctx, user.ID)
	if err != nil {
		return err
	}

	return nil
}
