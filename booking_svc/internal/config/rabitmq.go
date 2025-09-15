package config

import (
	"github.com/streadway/amqp"
	"log"
)

func ConnectRabbit(url string) (*amqp.Connection, *amqp.Channel) {
	log.Printf("Attempting to connect to RabbitMQ at: %s", url)

	var conn *amqp.Connection
	var err error

	if conn, err = amqp.Dial(url); err != nil {
		log.Fatalf("FATAL: RabbitMQ connection failed: %v", err)
	}
	log.Println("RabbitMQ connection established")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("FATAL: RabbitMQ channel creation failed: %v", err)
	}
	log.Println("RabbitMQ channel opened")

	return conn, ch
}
