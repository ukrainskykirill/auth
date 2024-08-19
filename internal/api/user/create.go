package user

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ukrainskykirill/auth/internal/converter"
	prError "github.com/ukrainskykirill/auth/internal/error"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *guser.CreateRequest) (*guser.CreateResponse, error) {
	userID, err := i.userService.Create(ctx, converter.ToUserInFromGUser(req))
	if err != nil {
		switch {
		case errors.Is(err, prError.ErrInvalidEmail):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, prError.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, prError.ErrPassword):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, prError.ErrNameNotUnique):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &guser.CreateResponse{
		Id: userID,
	}, nil
}
