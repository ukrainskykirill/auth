package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ukrainskykirill/auth/internal/api/user"
	"github.com/ukrainskykirill/auth/internal/model"
	"github.com/ukrainskykirill/auth/internal/service"
	serviceMocks "github.com/ukrainskykirill/auth/internal/service/mocks"
	desc "github.com/ukrainskykirill/auth/pkg/user_v1"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		userID = gofakeit.Int64()
		name   = gofakeit.Name()
		email  = gofakeit.Email()
		role   = desc.UserRole_USER

		serviceErr = fmt.Errorf("service err")

		req = &desc.GetRequest{
			Id: userID,
		}

		timestamp = timestamppb.Now()

		res = &desc.GetResponse{
			Id:        userID,
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: timestamp,
			UpdatedAt: timestamp,
		}

		User = &model.User{
			ID:        userID,
			Name:      name,
			Email:     email,
			Role:      role.String(),
			CreatedAt: timestamp.AsTime(),
			UpdatedAt: timestamp.AsTime(),
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, req.Id).Return(User, nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  status.Error(codes.Internal, serviceErr.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, req.Id).Return(&model.User{}, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			usersServiceMock := tt.userServiceMock(mc)

			api := user.NewImplementation(
				usersServiceMock,
			)

			chatID, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, chatID)
		})
	}

}
