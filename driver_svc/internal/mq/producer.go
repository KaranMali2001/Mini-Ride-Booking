package mq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type Producer struct {
	channel *amqp.Channel
}

type BookingAcceptedEvent struct {
	BookingID  string `json:"booking_id"`
	DriverID   string `json:"driver_id"`
	RideStatus string `json:"ride_status"`
}

func NewProducer(channel *amqp.Channel) *Producer {
	return &Producer{channel: channel}
}

func (p *Producer) PublishBookingAccepted(event BookingAcceptedEvent) error {
	// Declare the queue
	_, err := p.channel.QueueDeclare(
		"booking.accepted", // name
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Marshal the event to JSON
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Publish the message
	err = p.channel.Publish(
		"",                 // exchange
		"booking.accepted", // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		return err
	}

	log.Printf("Published booking.accepted event: %s", string(body))
	return nil
}
