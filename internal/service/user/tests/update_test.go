package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ukrainskykirill/auth/internal/cache"
	cacheMocks "github.com/ukrainskykirill/auth/internal/cache/mocks"
	prError "github.com/ukrainskykirill/auth/internal/error"
	"github.com/ukrainskykirill/auth/internal/model"
	"github.com/ukrainskykirill/auth/internal/repository"
	repoMocks "github.com/ukrainskykirill/auth/internal/repository/mocks"
	"github.com/ukrainskykirill/auth/internal/service/user"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepository
	type userCacheMockFunc func(mc *minimock.Controller) cache.UserCache

	type args struct {
		ctx context.Context
		req *model.UserInUpdate
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		repoErr    = fmt.Errorf("repo error")
		serviceErr = fmt.Errorf("service.Update - %w", prError.ErrInvalidEmail)

		userID = gofakeit.Int64()
		name   = gofakeit.Name()
		email  = gofakeit.Email()

		UserInUpdate = &model.UserInUpdate{
			Name:  name,
			Email: email,
		}
		UserInUpdateInvalidEmail = &model.UserInUpdate{
			ID:    userID,
			Name:  name,
			Email: name,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name         string
		args         args
		err          error
		userRepoMock userRepoMockFunc
		userCache    userCacheMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: UserInUpdate,
			},
			err: nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, UserInUpdate).Return(nil)
				return mock
			},
			userCache: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, UserInUpdate.ID).Return(nil)
				return mock
			},
		},
		{
			name: "error invalid email",
			args: args{
				ctx: ctx,
				req: UserInUpdateInvalidEmail,
			},
			err: serviceErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				return mock
			},
			userCache: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: UserInUpdate,
			},
			err: repoErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, UserInUpdate).Return(repoErr)
				return mock
			},
			userCache: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			usersRepo := tt.userRepoMock(mc)
			cacheUser := tt.userCache(mc)

			service := user.NewServ(
				usersRepo, cacheUser,
			)

			err := service.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
