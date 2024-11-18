package user_create

import (
	"context"

	"github.com/ukrainskykirill/auth/internal/client/rabbitmq"
	"github.com/ukrainskykirill/auth/internal/repository"
	consumerService "github.com/ukrainskykirill/auth/internal/service/consumer"
)

type userCreateService struct {
	repo     repository.UserRepository
	consumer rabbitmq.IConsumer
}

func NewUserCreateService(userRepo repository.UserRepository, consumer rabbitmq.IConsumer) consumerService.ConsumerService {
	return &userCreateService{
		repo:     userRepo,
		consumer: consumer,
	}
}

func (c *userCreateService) RunConsumer(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-c.run(ctx):
			if err != nil {
				return err
			}
		}
	}
}

func (s *userCreateService) run(ctx context.Context) <-chan error {
	errChan := make(chan error)

	go func() {
		defer close(errChan)

		errChan <- s.consumer.Consume(ctx, s.UserCreateHandler)
	}()

	return errChan
}
