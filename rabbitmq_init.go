package gutils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var RabbitConn *amqp.Connection

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func RabbitMQInit(amqpURL string) {
	var err error
	RabbitConn, err = amqp.Dial(amqpURL)
	failOnError(err, "Failed to connect to RabbitMQ")
}
