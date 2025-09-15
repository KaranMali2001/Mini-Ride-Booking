package mq

import (
	"encoding/json"

	"github.com/KaranMali2001/mini-ride-booking/common/logger"
	"github.com/streadway/amqp"
)

type Producer struct {
	channel *amqp.Channel
}

type BookingCreatedEvent struct {
	BookingID  string  `json:"booking_id"`
	PickupLoc  LatLng  `json:"pickuploc"`
	Dropoff    LatLng  `json:"dropoff"`
	Price      float64 `json:"price"`
	RideStatus string  `json:"ride_status"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func NewProducer(channel *amqp.Channel) *Producer {
	return &Producer{channel: channel}
}

func (p *Producer) PublishBookingCreated(event BookingCreatedEvent) error {
	// Declare the queue
	_, err := p.channel.QueueDeclare(
		"booking.created", // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
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
		"",                // exchange
		"booking.created", // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		return err
	}

	logger.Infof("Published booking.created event: %s", string(body))
	return nil
}
