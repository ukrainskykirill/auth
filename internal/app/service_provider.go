package app

import (
	"context"
	"log"

	"github.com/ukrainskykirill/auth/internal/api/user"
	userApi "github.com/ukrainskykirill/auth/internal/api/user"
	"github.com/ukrainskykirill/auth/internal/client/db"
	"github.com/ukrainskykirill/auth/internal/client/db/pg"
	"github.com/ukrainskykirill/auth/internal/client/db/transaction"
	"github.com/ukrainskykirill/auth/internal/closer"
	"github.com/ukrainskykirill/auth/internal/config"
	"github.com/ukrainskykirill/auth/internal/repository"
	userRepo "github.com/ukrainskykirill/auth/internal/repository/user"
	"github.com/ukrainskykirill/auth/internal/service"
	userServ "github.com/ukrainskykirill/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient  db.Client
	txManager db.TxManager

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

func (sp *serviceProvider) DBClient(ctx context.Context) db.Client {
	if sp.dbClient == nil {
		cl, err := pg.New(ctx, sp.PGConfig().URL())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		closer.Add(cl.Close)

		sp.dbClient = cl
	}

	return sp.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (sp *serviceProvider) UserRepo(ctx context.Context) repository.UserRepository {
	if sp.userRepo == nil {
		sp.userRepo = userRepo.NewUserRepository(sp.DBClient(ctx))
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
