package user

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	prError "github.com/ukrainskykirill/auth/internal/error"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *guser.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		switch {
		case errors.Is(err, prError.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &emptypb.Empty{}, nil
}
