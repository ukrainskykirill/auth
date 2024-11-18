package rabbitmq

import (
	"context"

	"github.com/ukrainskykirill/auth/internal/client/rabbitmq/consumer"
)

type IConsumer interface {
	Consume(ctx context.Context, queue string, handler consumer.MsgHandler) error
}
