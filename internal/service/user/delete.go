package user

import (
	"context"
	"fmt"
	prError "github.com/ukrainskykirill/auth/internal/error"
)

func (s *userServ) Delete(ctx context.Context, userID int64) error {
	isExist, err := s.repo.IsExistByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("service.Delete - %w", err)
	}
	if !isExist {
		return fmt.Errorf("service.Delete - %w", prError.ErrUserNotFound)
	}

	err = s.repo.Delete(ctx, userID)
	if err != nil {
		return fmt.Errorf("service.Delete - %w", err)
	}
	return nil
}
