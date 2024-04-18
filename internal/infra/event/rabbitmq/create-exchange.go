package rabbitmq

import (
	"fmt"

	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/logs"
	"github.com/streadway/amqp"
)

func CreateExchangeAndQueue(rabbitMQChannel *amqp.Channel, exchangeName, exchangeType, queueName, routingKeyName string) {
	err := rabbitMQChannel.ExchangeDeclare(
		exchangeName, // exchange name
		exchangeType, // exchange type
		true,         // durable
		false,        // deletable when not used
		false,        // exclusive (deleted when channel conection closes)
		false,        // no-wait
		nil,          // additional arguments
	)
	logs.FailOnError(err, "Error declaring exchange")

	queue, err := rabbitMQChannel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // deletable when not used
		false,     // exclusive (deleted when channel conection closes)
		false,     // no-wait
		nil,       // additional arguments
	)
	logs.FailOnError(err, "Error declaring queue")

	err = rabbitMQChannel.QueueBind(
		queue.Name,     // queue name
		routingKeyName, // routing key
		exchangeName,   // exchange name
		false,          // no-wait
		nil,            // additional arguments
	)
	logs.FailOnError(err, "Error on binding queue with exchange")

	fmt.Println("Exchange, queue and bind created successfully!")
}
