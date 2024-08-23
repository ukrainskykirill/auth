package user

import (
	"context"
	"fmt"
)

func (s *userServ) Delete(ctx context.Context, userID int64) error {
	err := s.repo.Delete(ctx, userID)
	if err != nil {
		return fmt.Errorf("service.Delete - %w", err)
	}
	return nil
}
