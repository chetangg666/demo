package rabbitmqq

import (
	"log"

	"github.com/streadway/amqp"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func Config() (ch *amqp.Channel, q amqp.Queue, conn *amqp.Connection) {
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@rabbit1:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	ch, err = conn.Channel()
	FailOnError(err, "Failed to open a channel")

	q, err = ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	FailOnError(err, "Failed to declare a queue")
	return ch, q, conn
}
