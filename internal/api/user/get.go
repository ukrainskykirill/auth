package user

import (
	"context"
	"github.com/ukrainskykirill/auth/internal/converter"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Get(ctx context.Context, req *guser.GetRequest) (*guser.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return converter.ToUserFromService(user), nil
}
