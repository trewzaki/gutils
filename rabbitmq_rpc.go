package gutils

import (
	"log"

	"github.com/streadway/amqp"
)

func ServerRPC(exchangeName string, topicName string, queueName string, fn func(amqp.Delivery) []byte) {
	ch, err := RabbitConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a exchange")

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp.Table{
			"x-message-ttl": int32(0),
		}, // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,       // queue name
		topicName,    // routing key
		exchangeName, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			go func(d amqp.Delivery) {
				result := fn(d)
				err = ch.Publish(
					"",        // exchange
					d.ReplyTo, // routing key
					false,     // mandatory
					false,     // immediate
					amqp.Publishing{
						ContentType:   "application/json",
						CorrelationId: d.CorrelationId,
						Body:          result,
					})
				failOnError(err, "Failed to publish a message")
			}(d)
		}
	}()
	log.Printf(" [server_rpc] running exchange=%s, topic=%s, queue=%s", exchangeName, topicName, queueName)
	<-forever
}

func ClientRPC(message []byte, exchangeName string, topicName string, corrID string) []byte {

	ch, err := RabbitConn.Channel()
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a exchange")

	consumer, err := ch.Consume(
		"amq.rabbitmq.reply-to", // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-waitx``
		nil,   // args
	)
	failOnError(err, "Failed to publish a message")

	err = ch.Publish(
		exchangeName, // exchange
		topicName,    // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrID,
			ReplyTo:       "amq.rabbitmq.reply-to",
			Body:          message,
		})
	failOnError(err, "Failed to publish a message")

	// fmt.Printf(" [x] Request: %s\n", message)

	for d := range consumer {
		if corrID == d.CorrelationId {

			return d.Body
		}
		// log.Printf(" [x] %s", d.Body)
	}

	// log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	return nil
}
