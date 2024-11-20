package interceptor

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/ukrainskykirill/auth/internal/model"
	"github.com/ukrainskykirill/auth/internal/service"
)

var availableEndpoints = []string{
	"/auth_v1.AuthV1/GetTokens",
	"/auth_v1.AuthV1/Login",
	"/user_v1.UserV1/Create",
}

type interceptor struct {
	authService service.AuthService
}

type validator interface {
	Validate() error
}

func NewInterceptor(authService service.AuthService) interceptor {
	return interceptor{
		authService: authService,
	}
}

func (i *interceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(
				codes.NotFound, fmt.Errorf("not found incoming context, %v", md).Error(),
			)
		}

		if !slices.Contains(availableEndpoints, info.FullMethod) {
			userClaims, err := i.authentication(ctx, md)
			if err != nil {
				return nil, status.Error(
					codes.NotFound, fmt.Errorf("failed authentication: %v", err).Error(),
				)
			}

			err = i.authorization(ctx, userClaims.Role, info.FullMethod)
			if err != nil {
				return nil, status.Error(
					codes.NotFound, fmt.Errorf("failed authentication: %v", err).Error(),
				)
			}
		}

		if val, ok := req.(validator); ok {
			if err := val.Validate(); err != nil {
				return nil, err
			}
		}

		return handler(ctx, req)
	}
}

func (i *interceptor) authentication(ctx context.Context, md metadata.MD) (*model.UserAccessClaims, error) {

	bearerToken := md["authorization"]

	if !strings.HasPrefix(bearerToken[0], "Bearer") {
		return &model.UserAccessClaims{}, fmt.Errorf("token doesnt have Bearer prefix: %s", bearerToken)
	}

	splitedTokens := strings.Split(bearerToken[0], " ")
	fmt.Println(splitedTokens[1])

	return i.authService.ParseAccessToken(splitedTokens[1])
}

func (i *interceptor) authorization(ctx context.Context, role string, endpoint string) error {
	err := i.authService.Check(ctx, role, endpoint)
	if err != nil {
		return err
	}
	return nil
}
