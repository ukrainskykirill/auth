package user

import (
	"context"
	"github.com/ukrainskykirill/auth/internal/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userServ) Create(ctx context.Context, user *model.UserIn) (int64, error) {
	err := validateEmail(user.Email)
	if err != nil {
		return 0, err
	}

	err = validatePassword(user.Password, user.PasswordConfirm)
	if err != nil {
		return 0, err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, status.Error(codes.InvalidArgument, err.Error())
	}
	user.Password = string(password)

	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
