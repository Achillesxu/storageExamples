// Package rabbitmq
// Time    : 2022/6/27 22:38
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package rabbitmq

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"testing"
)

const host = "amqp://root:123456@192.168.202.158:5672"

func TestPublish(t *testing.T) {
	q := New(host)
	defer q.Close()
	q.Bind("test")

	q2 := New(host)
	defer q2.Close()
	q2.Bind("test")

	q3 := New(host)
	defer q3.Close()

	expect := "test"
	q3.Publish("test2", "any")
	q3.Publish("test", expect)

	c := q.Consume()
	msg := <-c
	var actual interface{}
	err := json.Unmarshal(msg.Body, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expect {
		t.Errorf("expected %s, actual %s", expect, actual)
	}
	if msg.ReplyTo != q3.Name {
		t.Error(msg)
	}

	c2 := q2.Consume()
	msg = <-c2
	err = json.Unmarshal(msg.Body, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expect {
		t.Errorf("expected %s, actual %s", expect, actual)
	}
	if msg.ReplyTo != q3.Name {
		t.Error(msg)
	}
	q2.Send(msg.ReplyTo, "test3")
	c3 := q3.Consume()
	msg = <-c3
	if string(msg.Body) != `"test3"` {
		t.Error(string(msg.Body))
	}
}

func TestSend(t *testing.T) {
	q := New(host)
	defer q.Close()

	q2 := New(host)
	defer q2.Close()

	expect := "test"
	expect2 := "test2"
	q2.Send(q.Name, expect)
	q2.Send(q2.Name, expect2)

	c := q.Consume()
	msg := <-c
	var actual interface{}
	err := json.Unmarshal(msg.Body, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expect {
		t.Errorf("expected %s, actual %s", expect, actual)
	}

	c2 := q2.Consume()
	msg = <-c2
	err = json.Unmarshal(msg.Body, &actual)
	if err != nil {
		t.Error(err)
	}
	if actual != expect2 {
		t.Errorf("expected %s, actual %s", expect2, actual)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", err, msg)
	}
}

func sendHello() {
	conn, err := amqp.Dial(host)
	failOnError(err, "send Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			failOnError(err, "send Failed to close RabbitMQ")
		}
	}(conn)

	ch, err := conn.Channel()
	failOnError(err, "send Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			failOnError(err, "send Failed to close a channel")
		}
	}(ch)

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "send Failed to declare a queue")
	body := "Hello World"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte("hell01")})
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte("hell01")})
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte("hell01")})
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte("hell01")})

	failOnError(err, "send Failed to publish a msg")
	log.Printf("[x] Sent %s\n", body)
}

func recvHello() {
	conn, err := amqp.Dial(host)
	failOnError(err, "recv Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			failOnError(err, "recv Failed to close RabbitMQ")
		}
	}(conn)

	ch, err := conn.Channel()
	failOnError(err, "recv Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			failOnError(err, "recv Failed to close a channel")
		}
	}(ch)

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "recv Failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "recv Failed to register a consumer")
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Recv a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func TestHelloSendRecv(t *testing.T) {
	sendHello()
	recvHello()
}
