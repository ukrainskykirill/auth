package consumer

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MsgHandler func(ctx context.Context, msg *amqp.Delivery) error

type Consumer struct {
	connection    *amqp.Connection
	channel       *amqp.Channel
	done          chan error
	handler       MsgHandler
	tag           string
	queue         string
	maxRetryCount int
}

func NewConsumer(url, queue string, maxRetryCount int) (*Consumer, error) {
	connection, err := amqp.DialConfig(
		url,
		amqp.Config{
			Heartbeat: time.Duration(60 * time.Minute),
		},
	)
	if err != nil {
		return &Consumer{}, fmt.Errorf("consumer dial: %s", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		return &Consumer{}, fmt.Errorf("consumer channel: %s", err)
	}

	return &Consumer{
		connection:    connection,
		channel:       channel,
		done:          make(chan error),
		handler:       nil,
		tag:           "",
		queue:         queue,
		maxRetryCount: maxRetryCount,
	}, nil
}

func (c *Consumer) Consume(ctx context.Context, handler MsgHandler) error {
	c.handler = handler

	deliveries, err := c.channel.Consume(
		c.queue,
		c.tag,
		false,
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
	for msg := range deliveries {
		fmt.Printf(
			"got %dB delivery: [%v] %q",
			len(msg.Body),
			msg.DeliveryTag,
			msg.Body,
		)
		if err := c.handler(ctx, &msg); err != nil {
			c.msgRetry(msg)
		} else {
			msg.Ack(false)
		}
	}
	fmt.Printf("handle: deliveries channel closed")
	done <- nil
}

func (c *Consumer) msgRetry(msg amqp.Delivery) {
	if msg.Headers["x-death"] != nil {
		for _, death := range msg.Headers["x-death"].([]interface{}) {
			deathMap := death.(amqp.Table)
			if deathMap["reason"] == "rejected" {
				count, ok := deathMap["count"].(int)
				if ok && count < c.maxRetryCount {
					msg.Nack(false, false)
				} else {
					msg.Ack(false)
					fmt.Printf("maximum retries has been exceeded: %s", msg.MessageId)
				}
				break
			}
		}
	} else {
		msg.Nack(false, false)
	}
}
