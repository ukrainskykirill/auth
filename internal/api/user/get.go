package user

import (
	"context"
	"errors"
	"github.com/ukrainskykirill/auth/internal/converter"
	prError "github.com/ukrainskykirill/auth/internal/error"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Get(ctx context.Context, req *guser.GetRequest) (*guser.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		switch {
		case errors.Is(err, prError.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return converter.ToUserFromService(user), nil
}
