package rabbitmq

import (
	"context"

	"github.com/ukrainskykirill/auth/internal/client/rabbitmq/consumer"
)

type IConsumer interface {
	Consume(ctx context.Context, handler consumer.MsgHandler) error
}
