package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	prError "github.com/ukrainskykirill/auth/internal/error"
	"github.com/ukrainskykirill/auth/internal/model"
	"github.com/ukrainskykirill/auth/internal/repository"
	repoMocks "github.com/ukrainskykirill/auth/internal/repository/mocks"
	"github.com/ukrainskykirill/auth/internal/service/user"
	desc "github.com/ukrainskykirill/auth/pkg/user_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req *model.UserIn
	}

	var (
		ctx                 = context.Background()
		mc                  = minimock.NewController(t)
		repoErr             = fmt.Errorf("repo error")
		invalidEmailErr     = fmt.Errorf("service.Create - %w", prError.ErrInvalidEmail)
		validatePasswordErr = fmt.Errorf("service.Create - %w", fmt.Errorf("%w: passwords doesnt match", prError.ErrPassword))

		userID   = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, 1)
		role     = desc.UserRole_USER

		UserIn = &model.UserIn{
			Name:            name,
			Email:           email,
			Role:            role.String(),
			Password:        password,
			PasswordConfirm: password,
		}
		UserInvalidEmail = &model.UserIn{
			Name:            name,
			Email:           name,
			Role:            role.String(),
			Password:        password,
			PasswordConfirm: password,
		}
		UserInvalidPassword = &model.UserIn{
			Name:            name,
			Email:           email,
			Role:            role.String(),
			Password:        password,
			PasswordConfirm: name,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name         string
		args         args
		want         int64
		err          error
		userRepoMock userRepoMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: UserIn,
			},
			want: userID,
			err:  nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, UserIn).Return(userID, nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: UserIn,
			},
			want: 0,
			err:  repoErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, UserIn).Return(0, repoErr)
				return mock
			},
		},
		{
			name: "invalid email error",
			args: args{
				ctx: ctx,
				req: UserInvalidEmail,
			},
			want: 0,
			err:  invalidEmailErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				return mock
			},
		},
		{
			name: "validate password error",
			args: args{
				ctx: ctx,
				req: UserInvalidPassword,
			},
			want: 0,
			err:  validatePasswordErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			usersRepo := tt.userRepoMock(mc)

			service := user.NewServ(
				usersRepo,
			)

			id, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, id)
		})
	}
}
