package config

import (
	"fmt"
	"os"
)

const (
	rabbitMQHostEnv  = "RABBITMQ_HOST"
	rabbitMQPortEnv  = "RABBITMQ_PORT"
	rabbitMQUserEnv  = "RABBITMQ_USER"
	rabbitMQPassEnv  = "RABBITMQ_PASS"
	rabbitMQQueueEnv = "RABBITMQ_QUEUE"
)

type RabbitMQConsumerConfig interface {
	URL() string
	Queue() string
}

type rabbitMQConsumerConfig struct {
	host     string
	port     string
	username string
	password string
	queue    string
}

func NewRabbitMQConsumerConfig() (RabbitMQConsumerConfig, error) {
	host, ok := os.LookupEnv(rabbitMQHostEnv)
	if !ok {
		return &rabbitMQConsumerConfig{}, errVariableNotFound
	}

	port, ok := os.LookupEnv(rabbitMQPortEnv)
	if !ok {
		return &rabbitMQConsumerConfig{}, errVariableNotFound
	}

	username, ok := os.LookupEnv(rabbitMQUserEnv)
	if !ok {
		return &rabbitMQConsumerConfig{}, errVariableNotFound
	}

	password, ok := os.LookupEnv(rabbitMQPassEnv)
	if !ok {
		return &rabbitMQConsumerConfig{}, errVariableNotFound
	}

	queue, ok := os.LookupEnv(rabbitMQQueueEnv)
	if !ok {
		return &rabbitMQConsumerConfig{}, errVariableNotFound
	}

	return &rabbitMQConsumerConfig{
		host:     host,
		port:     port,
		username: username,
		password: password,
		queue:    queue,
	}, nil
}

func (c *rabbitMQConsumerConfig) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", c.username, c.password, c.host, c.port)
}

func (c *rabbitMQConsumerConfig) Queue() string {
	return c.queue
}
