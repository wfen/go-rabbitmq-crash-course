package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Service -
type Service interface {
	Connect() error
	Publish(message string) error
	Consume()
}

// RabbitMQ -
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// Connect - establishes a connection to our RabbitMQ instance
// and declares the queue we are going to be using
func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ")
	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to RabbitMQ")

	// We need to open a channel over our AMQP connection
	// This allows us to declare queues & subsequently consume/publish messages
	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err
	}

	// Here we declare our new queue that we want to publish to and consume from
	_, err = r.Channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	return nil
}

// Publish - takes in a string message and publishes to a queue
func (r *RabbitMQ) Publish(message string) error {
	err := r.Channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

// Consume - consumes messages from our test queue
func (r *RabbitMQ) Consume() {
	msgs, err := r.Channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	for msg := range msgs {
		fmt.Printf("Received message: %s\n", msg.Body)
	}
}

// NewRabbitMQService - returns a pointer to a new RabbitMQ service
func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
