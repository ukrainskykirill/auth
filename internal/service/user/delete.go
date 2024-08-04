package user

import (
	"context"
)

func (s *userServ) Delete(ctx context.Context, userID int64) error {
	err := s.repo.Delete(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
