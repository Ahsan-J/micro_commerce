package rabbitmq

import (
	"cart/helper"
	"cart/rabbitmq/cart"
	"fmt"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection

// ConnectToRabbitMQ establishes a connection and returns channel
func ConnectToRabbitMQ() *amqp.Connection {
	config := helper.GetConfiguration()

	conn, err := amqp.Dial(config.RabbitURL)
	helper.FailOnError(err, "Failed to connect to RabbitMQ")
	connection = conn
	fmt.Println("Connected to RabbitMQ at ", config.RabbitURL)
	cart.RegisterCartQueue(GetChannel())
	// defer connection.Close()
	return conn
}

// GetChannel exports the created channel
func GetChannel() *amqp.Channel {
	if connection.IsClosed() {
		ConnectToRabbitMQ()
	}
	ch, err := connection.Channel()

	helper.FailOnError(err, "Failed to open a channel")
	// defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	helper.FailOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	helper.FailOnError(err, "Failed to publish a message")
	return ch
}
