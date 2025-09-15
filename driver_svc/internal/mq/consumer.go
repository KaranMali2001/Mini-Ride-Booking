package mq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/db/generated"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/streadway/amqp"
)

type Consumer struct {
	channel *amqp.Channel
	queries *generated.Queries
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

func NewConsumer(channel *amqp.Channel, queries *generated.Queries) *Consumer {
	return &Consumer{
		channel: channel,
		queries: queries,
	}
}

func (c *Consumer) StartConsumingBookingCreated(ctx context.Context) error {
	// Declare the queue
	_, err := c.channel.QueueDeclare(
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

	// Start consuming messages
	msgs, err := c.channel.Consume(
		"booking.created", // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		return err
	}

	log.Printf("Waiting for booking.created messages...")

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("Consumer context cancelled, stopping...")
				return
			case msg := <-msgs:
				if len(msg.Body) == 0 {
					continue
				}

				var event BookingCreatedEvent
				if err := json.Unmarshal(msg.Body, &event); err != nil {
					log.Printf("Failed to unmarshal booking.created event: %v", err)
					continue
				}

				log.Printf("Received booking.created event: %+v", event)

				// Create a job entry for drivers to see
				if err := c.handleBookingCreated(ctx, event); err != nil {
					log.Printf("Failed to handle booking.created event: %v", err)
				}
			}
		}
	}()

	return nil
}

func (c *Consumer) handleBookingCreated(ctx context.Context, event BookingCreatedEvent) error {
	// Parse booking ID from string to UUID
	var bookingUUID pgtype.UUID
	err := bookingUUID.Scan(event.BookingID)
	if err != nil {
		log.Printf("Failed to parse booking ID: %v", err)
		return err
	}

	// Create a job entry that drivers can see and accept
	_, err = c.queries.CreateJob(ctx, generated.CreateJobParams{
		BookingID:    bookingUUID,
		DriverID:     pgtype.UUID{}, // NULL initially
		RideStatus:   event.RideStatus,
		PickuplocLat: event.PickupLoc.Lat,
		PickuplocLng: event.PickupLoc.Lng,
		DropoffLat:   event.Dropoff.Lat,
		DropoffLng:   event.Dropoff.Lng,
		Price:        int32(event.Price),
	})

	if err != nil {
		log.Printf("Failed to create job for booking: %v", err)
		return err
	}

	log.Printf("Successfully created job for booking %s", event.BookingID)
	return nil
}