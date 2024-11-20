package user

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	prError "github.com/ukrainskykirill/auth/internal/error"
	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *userServ) Create(ctx context.Context, user *model.UserIn) (int64, error) {
	err := validateEmail(user.Email)
	if err != nil {
		return 0, fmt.Errorf("service.Create - %w", err)
	}

	err = validatePassword(user.Password, user.PasswordConfirm)
	if err != nil {
		return 0, fmt.Errorf("service.Create - %w", err)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("service.Create - %w: generate password error %w", prError.ErrPassword, err)
	}

	user.Password = string(password)

	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
