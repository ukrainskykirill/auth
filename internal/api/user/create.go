package user

import (
	"context"
	"github.com/ukrainskykirill/auth/internal/converter"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Create(ctx context.Context, req *guser.CreateRequest) (*guser.CreateResponse, error) {
	userID, err := i.userService.Create(ctx, converter.ToUserInFromGUser(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &guser.CreateResponse{
		Id: userID,
	}, nil
}
