package utils

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/faraji-fuji/miniature-umbrella/src/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func SendToExchange(notification models.Notification) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"miniature_umbrella", // name
		"direct",             // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		failOnError(err, "Failed to marshal notification")
	}

	err = ch.PublishWithContext(ctx,
		"miniature_umbrella", // exchange
		notification.Channel, // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        notificationJSON,
		})

	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", notificationJSON)
}
