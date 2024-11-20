package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ukrainskykirill/auth/internal/model"
	gauth "github.com/ukrainskykirill/auth/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *gauth.LoginRequest) (*gauth.TokensResponse, error) {
	tokens, err := i.authService.Login(ctx, &model.LoginIn{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &gauth.TokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
