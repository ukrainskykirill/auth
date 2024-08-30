package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ukrainskykirill/auth/internal/cache"
	cacheMocks "github.com/ukrainskykirill/auth/internal/cache/mocks"
	"github.com/ukrainskykirill/auth/internal/model"
	"github.com/ukrainskykirill/auth/internal/repository"
	repoMocks "github.com/ukrainskykirill/auth/internal/repository/mocks"
	"github.com/ukrainskykirill/auth/internal/service/user"
	desc "github.com/ukrainskykirill/auth/pkg/user_v1"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepository
	type userCacheMockFunc func(mc *minimock.Controller) cache.UserCache

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx      = context.Background()
		mc       = minimock.NewController(t)
		repoErr  = fmt.Errorf("repo error")
		cacheErr = fmt.Errorf("cache error")

		userID = gofakeit.Int64()
		name   = gofakeit.Name()
		email  = gofakeit.Email()
		role   = desc.UserRole_USER

		UserIn = &model.User{
			Name:      name,
			Email:     email,
			Role:      role.String(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name         string
		args         args
		want         *model.User
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
			want: UserIn,
			err:  nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(UserIn, nil)
				return mock
			},
			userCache: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(nil, nil)
				mock.CreateMock.Expect(ctx, UserIn).Return(nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: userID,
			},
			want: &model.User{},
			err:  repoErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(&model.User{}, repoErr)
				return mock
			},
			userCache: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(nil, nil)
				return mock
			},
		},
		{
			name: "success cache case",
			args: args{
				ctx: ctx,
				req: userID,
			},
			want: UserIn,
			err:  nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				return mock
			},
			userCache: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(UserIn, nil)
				return mock
			},
		},
		{
			name: "cache error case",
			args: args{
				ctx: ctx,
				req: userID,
			},
			want: &model.User{},
			err:  cacheErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				return mock
			},
			userCache: func(mc *minimock.Controller) cache.UserCache {
				mock := cacheMocks.NewUserCacheMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(nil, cacheErr)
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

			id, err := service.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, id)
		})
	}
}
