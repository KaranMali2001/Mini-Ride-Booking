package config

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnectRabbit(url string) (*amqp.Connection, *amqp.Channel) {

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("FATAL: RabbitMQ connection failed: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("FATAL: RabbitMQ channel creation failed: %v", err)
	}
	log.Println("RabbitMQ channel opened")

	return conn, ch
}
