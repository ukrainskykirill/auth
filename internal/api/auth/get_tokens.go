package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	gauth "github.com/ukrainskykirill/auth/pkg/auth_v1"
)

func (i *Implementation) GetTokens(ctx context.Context, req *gauth.GetTokensRequest) (*gauth.TokensResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(md, ok)

	tokens, err := i.authService.GetTokens(ctx, req.OldRefreshToken)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &gauth.TokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
