package user_create

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/ukrainskykirill/auth/internal/model"
)

func (s *userCreateService) UserCreateHandler(ctx context.Context, msg *amqp.Delivery) error {
	userIn := &model.UserIn{}
	err := json.Unmarshal(msg.Body, userIn)
	if err != nil {
		return err
	}

	userID, err := s.repo.Create(ctx, userIn)
	if err != nil {
		return err
	}

	log.Printf("User with id %d created\n", userID)

	return nil
}
