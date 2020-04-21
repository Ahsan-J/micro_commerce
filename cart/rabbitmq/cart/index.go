package cart

import (
	"cart/helper"
	"cart/modal"
	"encoding/json"

	"github.com/streadway/amqp"
)

var events = []string{"hello"}

var cartChannel *amqp.Channel

// RegisterCartQueue registers the cart queue and exchange with all of its routing keys
func RegisterCartQueue(ch *amqp.Channel) {
	config := helper.GetConfiguration()
	err := ch.ExchangeDeclare(
		config.Exchange.Cart, // name
		"topic",              // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)

	helper.FailOnError(err, "Cannot create an exchange")

	q, err := ch.QueueDeclare(
		config.Queue.Cart, // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)

	helper.FailOnError(err, "Cannot create a Queue")

	// Binding Routing key to Queue to Exhange
	for _, event := range events {
		ch.QueueBind(
			q.Name,               // queue name
			event,                // routing key
			config.Exchange.Cart, // exchange
			false,                // no-wait
			nil,                  // arguments
		)
	}

	cartChannel = ch
	Consumer(ch, q)
}

func publish(routingKey string, data interface{}) {
	sendingObjet := modal.MessageQueueObject{
		RoutingKey: "hello",
		Data:       data,
	}
	bytes, err := json.Marshal(sendingObjet)
	helper.FailOnError(err, "Error while marshalling data to SendHelloWorld")

	config := helper.GetConfiguration()
	err = cartChannel.Publish(
		config.Exchange.Cart, // exchange
		routingKey,           // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})
	helper.FailOnError(err, "Failed to publish a message")
}
