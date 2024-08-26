package user

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ukrainskykirill/auth/internal/converter"
	prError "github.com/ukrainskykirill/auth/internal/error"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
)

func (i *Implementation) Update(ctx context.Context, req *guser.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, converter.ToUserInUpdateFromGUser(req))
	if err != nil {
		switch {
		case errors.Is(err, prError.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, prError.ErrInvalidEmail):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &emptypb.Empty{}, nil
}
