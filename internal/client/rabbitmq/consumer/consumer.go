package consumer

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MsgHandler func(ctx context.Context, msg *amqp.Delivery) error

type Consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	done       chan error
	handler    MsgHandler
	tag        string
}

func NewConsumer(url string) (*Consumer, error) {
	connection, err := amqp.Dial(url)
	if err != nil {
		return &Consumer{}, fmt.Errorf("consumer dial: %s", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		return &Consumer{}, fmt.Errorf("consumer channel: %s", err)
	}

	return &Consumer{
		connection: connection,
		channel:    channel,
		done:       make(chan error),
		handler:    nil,
		tag:        "",
	}, nil
}

func (c *Consumer) Consume(ctx context.Context, queue string, handler MsgHandler) error {
	c.handler = handler
	if c.handler == nil {
		return fmt.Errorf("Message handler is nil")
	}

	deliveries, err := c.channel.Consume(
		queue,
		c.tag,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("consume error: %s", err)
	}

	go c.handle(ctx, deliveries, c.done)
	return nil
}

func (c *Consumer) Shutdown() error {
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := c.connection.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer fmt.Printf("AMQP shutdown OK")

	return <-c.done
}

func (c *Consumer) handle(ctx context.Context, deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		fmt.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		c.handler(ctx, &d)
		// d.Ack(false)
	}
	fmt.Printf("handle: deliveries channel closed")
	done <- nil
}
