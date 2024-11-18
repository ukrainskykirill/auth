package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx := context.Background()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}

	defer func() {
		_ = conn.Close()
	}()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel. Error: %s", err)
	}

	defer func() {
		_ = ch.Close()
	}()

	userIn := map[string]string{"email": "test@test.com", "name": "Test3", "password": "pass", "password_confirm": "pass", "role": "USER"}
	body, err := json.Marshal(userIn)
	if err != nil {
		fmt.Println(err)
	}

	err = ch.PublishWithContext(ctx,
		"amq.direct", // exchange
		"key",        // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		log.Fatalf("failed to publish a message. Error: %s", err)
	}

	log.Printf(" [x] Sent %s\n", body)
}
