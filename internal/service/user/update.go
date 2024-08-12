package user

import (
	"context"
	"fmt"
	prError "github.com/ukrainskykirill/auth/internal/error"
	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *userServ) Update(ctx context.Context, user *model.UserInUpdate) error {
	isExist, err := s.repo.IsExistByID(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("service.Update - %w", err)
	}
	if !isExist {
		return fmt.Errorf("service.Update - %w", prError.ErrUserNotFound)
	}

	if len(user.Email) != 0 {
		err := validateEmail(user.Email)
		if err != nil {
			return fmt.Errorf("service.Update - %w", err)
		}
	}

	if len(user.Name) != 0 {
		isExist, err := s.repo.IsExistByName(ctx, user.Name)
		if err != nil {
			return fmt.Errorf("service.Update - %w", err)
		}
		if isExist {
			return fmt.Errorf("service.Update - %w", prError.ErrNameNotUnique)
		}
	}

	err = s.repo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
