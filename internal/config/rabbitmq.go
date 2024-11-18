package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	rabbitMQHostEnv           = "RABBITMQ_HOST"
	rabbitMQPortEnv           = "RABBITMQ_PORT"
	rabbitMQUserEnv           = "RABBITMQ_USER"
	rabbitMQPassEnv           = "RABBITMQ_PASS"
	rabbitMQQueueEnv          = "RABBITMQ_QUEUE"
	rabbitMQQMaxRetryCountEnv = "RABBITMQ_MAX_RETRY_COUNT"
)

type RabbitMQConsumerConfig interface {
	URL() string
	Queue() string
	MaxRetryCount() int
}

type rabbitMQConsumerConfig struct {
	host          string
	port          string
	username      string
	password      string
	queue         string
	maxRetryCount int
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

	maxRetryCountStr, ok := os.LookupEnv(rabbitMQQMaxRetryCountEnv)
	if !ok {
		return &rabbitMQConsumerConfig{}, errVariableNotFound
	}

	maxRetryCount, err := strconv.Atoi(maxRetryCountStr)
	if err != nil {
		return nil, errVariableParse
	}

	return &rabbitMQConsumerConfig{
		host:          host,
		port:          port,
		username:      username,
		password:      password,
		queue:         queue,
		maxRetryCount: maxRetryCount,
	}, nil
}

func (c *rabbitMQConsumerConfig) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", c.username, c.password, c.host, c.port)
}

func (c *rabbitMQConsumerConfig) Queue() string {
	return c.queue
}

func (c *rabbitMQConsumerConfig) MaxRetryCount() int {
	return c.maxRetryCount
}
