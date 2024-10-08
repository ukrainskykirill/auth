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
	"github.com/ukrainskykirill/auth/internal/repository"
	repoMocks "github.com/ukrainskykirill/auth/internal/repository/mocks"
	"github.com/ukrainskykirill/auth/internal/service/user"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepository
	type userCacheMockFunc func(mc *minimock.Controller) cache.UserCache

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		repoErr = fmt.Errorf("repo error")
		resErr  = fmt.Errorf("service.Delete - %w", repoErr)

		userID = gofakeit.Int64()
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
				req: userID,
			},
			err: nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, userID).Return(nil)
				return mock
			},
			userCache: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.DeleteMock.Expect(ctx, userID).Return(nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: userID,
			},
			err: resErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, userID).Return(repoErr)
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

			err := service.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
