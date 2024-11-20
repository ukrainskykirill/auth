package app

import (
	"context"
	"fmt"
	"log"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/ukrainskykirill/platform_common/pkg/cache"
	"github.com/ukrainskykirill/platform_common/pkg/cache/redis"
	"github.com/ukrainskykirill/platform_common/pkg/closer"
	"github.com/ukrainskykirill/platform_common/pkg/db"
	"github.com/ukrainskykirill/platform_common/pkg/db/pg"
	"github.com/ukrainskykirill/platform_common/pkg/db/transaction"

	accessApi "github.com/ukrainskykirill/auth/internal/api/access"
	authApi "github.com/ukrainskykirill/auth/internal/api/auth"
	"github.com/ukrainskykirill/auth/internal/api/user"
	userApi "github.com/ukrainskykirill/auth/internal/api/user"
	prCache "github.com/ukrainskykirill/auth/internal/cache"
	authCache "github.com/ukrainskykirill/auth/internal/cache/auth"
	userCache "github.com/ukrainskykirill/auth/internal/cache/user"
	"github.com/ukrainskykirill/auth/internal/client/rabbitmq"
	rabbitmqConsumer "github.com/ukrainskykirill/auth/internal/client/rabbitmq/consumer"
	"github.com/ukrainskykirill/auth/internal/config"
	"github.com/ukrainskykirill/auth/internal/repository"
	accessRepo "github.com/ukrainskykirill/auth/internal/repository/access"
	userRepo "github.com/ukrainskykirill/auth/internal/repository/user"
	"github.com/ukrainskykirill/auth/internal/service"
	authServ "github.com/ukrainskykirill/auth/internal/service/auth"
	consumerService "github.com/ukrainskykirill/auth/internal/service/consumer"
	userConsumer "github.com/ukrainskykirill/auth/internal/service/consumer/user"
	userServ "github.com/ukrainskykirill/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig               config.PGConfig
	grpcConfig             config.GRPCConfig
	redisConfig            config.RedisConfig
	rabbitmqConsumerConfig config.RabbitMQConsumerConfig
	authConfig             config.AuthConfig

	redisPool   redigo.Pool
	redisClient cache.Client

	dbClient  db.Client
	txManager db.TxManager

	rabbitMQConsumer rabbitmq.IConsumer

	userCache prCache.UserCache
	authCache prCache.AuthCache

	userRepo   repository.UserRepository
	accessRepo repository.AccessRepository

	userServ service.UserService
	authServ service.AuthService

	userCreateConsumer consumerService.ConsumerService

	userAPI   *user.Implementation
	authAPI   *authApi.Implementation
	accessAPI *accessApi.Implementation
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

func (sp *serviceProvider) RabbitMQConsumerConfig() config.RabbitMQConsumerConfig {
	if sp.rabbitmqConsumerConfig == nil {
		cfg, err := config.NewRabbitMQConsumerConfig()
		if err != nil {
			log.Fatalf("Error loading config: %s", err.Error())
		}
		sp.rabbitmqConsumerConfig = cfg
	}
	return sp.rabbitmqConsumerConfig
}

func (sp *serviceProvider) AuthConfig() config.AuthConfig {
	if sp.authConfig == nil {
		cfg, err := config.NewAuthConfig()
		if err != nil {
			log.Fatalf("Error loading config: %s", err.Error())
		}
		sp.authConfig = cfg
	}
	return sp.authConfig
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

func (sp *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if sp.txManager == nil {
		sp.txManager = transaction.NewTransactionManager(sp.DBClient(ctx).DB())
	}

	return sp.txManager
}

func (sp *serviceProvider) RedisConfig() config.RedisConfig {
	if sp.redisConfig == nil {
		cfg, err := config.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %s", err.Error())
		}

		sp.redisConfig = cfg
	}

	return sp.redisConfig
}

func (sp *serviceProvider) RedisPool(ctx context.Context) *redigo.Pool {
	sp.redisPool = redigo.Pool{
		MaxIdle:     sp.RedisConfig().MaxIdle(),
		IdleTimeout: sp.RedisConfig().IdleTimeout(),
		DialContext: func(ctx context.Context) (redigo.Conn, error) {
			return redigo.DialContext(ctx, "tcp", sp.RedisConfig().Address())
		},
	}
	return &sp.redisPool
}

func (sp *serviceProvider) RedisClient(ctx context.Context) cache.Client {
	if sp.redisClient == nil {
		sp.redisClient = redis.NewClient(sp.RedisPool(ctx), sp.RedisConfig().ConnectionTimeout())
	}
	return sp.redisClient
}

func (sp *serviceProvider) UserCache(ctx context.Context) prCache.UserCache {
	if sp.userCache == nil {
		sp.userCache = userCache.NewCache(sp.RedisClient(ctx))
	}
	return sp.userCache
}

func (sp *serviceProvider) AuthCache(ctx context.Context) prCache.AuthCache {
	if sp.authCache == nil {
		sp.authCache = authCache.NewCache(sp.RedisClient(ctx))
	}
	return sp.authCache
}

func (sp *serviceProvider) UserRepo(ctx context.Context) repository.UserRepository {
	if sp.userRepo == nil {
		sp.userRepo = userRepo.NewUserRepository(sp.DBClient(ctx))
	}

	return sp.userRepo
}

func (sp *serviceProvider) AccessRepo(ctx context.Context) repository.AccessRepository {
	if sp.accessRepo == nil {
		sp.accessRepo = accessRepo.NewAccessRepository(sp.DBClient(ctx))
	}

	return sp.accessRepo
}

func (sp *serviceProvider) UserService(ctx context.Context) service.UserService {
	if sp.userServ == nil {
		sp.userServ = userServ.NewServ(sp.UserRepo(ctx), sp.UserCache(ctx))
	}

	return sp.userServ
}

func (sp *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if sp.authServ == nil {
		sp.authServ = authServ.NewAuthServ(
			sp.UserRepo(ctx),
			sp.AccessRepo(ctx),
			sp.AuthCache(ctx),
			sp.AuthConfig().TokenSecret(),
			sp.AuthConfig().AccessTokenTTL(),
			sp.AuthConfig().RefreshTokenTTL(),
		)
	}
	return sp.authServ
}

func (sp *serviceProvider) UserAPI(ctx context.Context) *user.Implementation {
	if sp.userAPI == nil {
		sp.userAPI = userApi.NewImplementation(sp.UserService(ctx))
	}

	return sp.userAPI
}

func (sp *serviceProvider) AuthAPI(ctx context.Context) *authApi.Implementation {
	if sp.authAPI == nil {
		sp.authAPI = authApi.NewAuthImplementation(sp.AuthService(ctx))
	}

	return sp.authAPI
}

func (sp *serviceProvider) AccessAPI(ctx context.Context) *accessApi.Implementation {
	if sp.accessAPI == nil {
		sp.accessAPI = accessApi.NewAccessImplementation(sp.AuthService(ctx))
	}
	return sp.accessAPI
}

func (sp *serviceProvider) RabbitMQConsumer() rabbitmq.IConsumer {
	if sp.rabbitMQConsumer == nil {
		var err error
		sp.rabbitMQConsumer, err = rabbitmqConsumer.NewConsumer(
			sp.RabbitMQConsumerConfig().URL(),
			sp.RabbitMQConsumerConfig().Queue(),
			sp.RabbitMQConsumerConfig().MaxRetryCount(),
		)
		if err != nil {
			fmt.Println(err)
		}
	}

	return sp.rabbitMQConsumer
}

func (sp *serviceProvider) UserCreateConsumer(ctx context.Context) consumerService.ConsumerService {
	if sp.userCreateConsumer == nil {
		sp.userCreateConsumer = userConsumer.NewUserCreateService(
			sp.UserRepo(ctx),
			sp.RabbitMQConsumer(),
		)
	}

	return sp.userCreateConsumer
}
