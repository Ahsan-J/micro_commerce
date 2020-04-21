package cart

import (
	"cart/helper"
	"log"

	"github.com/streadway/amqp"
)

// Consumer is the consumption function
func Consumer(ch *amqp.Channel, q amqp.Queue) {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	helper.FailOnError(err, "Cannot create consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			log.Printf(" [x] %s", d.Body)

			d.Ack(true)
		}
	}()
	<-forever
}
