// Package rabbit
// Time    : 2022/7/4 21:29
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	conn, err := amqp.Dial(host)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func() { _ = conn.Close() }()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func() { _ = ch.Close() }()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

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
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
