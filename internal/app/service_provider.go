package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukrainskykirill/auth/internal/api/user"
	userApi "github.com/ukrainskykirill/auth/internal/api/user"
	"github.com/ukrainskykirill/auth/internal/closer"
	"github.com/ukrainskykirill/auth/internal/config"
	"github.com/ukrainskykirill/auth/internal/repository"
	userRepo "github.com/ukrainskykirill/auth/internal/repository/user"
	"github.com/ukrainskykirill/auth/internal/service"
	userServ "github.com/ukrainskykirill/auth/internal/service/user"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool   *pgxpool.Pool
	userRepo repository.UserRepository

	userServ service.UserService

	userApi *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) PGConfig() config.PGConfig {
	if sp.pgConfig == nil {
		cfg, err := config.NewDBConfig()
		if err != nil {
			log.Fatalf("Error loading config: %s", err.Error())
		}

		sp.pgConfig = cfg
	}

	return sp.pgConfig
}

func (sp *serviceProvider) GRPCConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("Error loading config: %s", err.Error())
		}
		sp.grpcConfig = cfg
	}
	return sp.grpcConfig
}

func (sp *serviceProvider) PGPool(ctx context.Context) *pgxpool.Pool {
	if sp.pgPool == nil {
		pool, err := pgxpool.New(ctx, sp.PGConfig().URL())
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v\n", err)
		}

		closer.Add(func() error {
			pool.Close()
			return nil
		})

		sp.pgPool = pool
	}

	return sp.pgPool
}

func (sp *serviceProvider) UserRepo(ctx context.Context) repository.UserRepository {
	if sp.userRepo == nil {
		sp.userRepo = userRepo.NewUserRepository(sp.PGPool(ctx))
	}

	return sp.userRepo
}

func (sp *serviceProvider) UserService(ctx context.Context) service.UserService {
	if sp.userServ == nil {
		sp.userServ = userServ.NewServ(sp.UserRepo(ctx))
	}

	return sp.userServ
}

func (sp *serviceProvider) UserAPI(ctx context.Context) *user.Implementation {
	if sp.userApi == nil {
		sp.userApi = userApi.NewImplementation(sp.UserService(ctx))
	}

	return sp.userApi
}
