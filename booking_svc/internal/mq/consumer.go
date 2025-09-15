package mq

import (
	"context"
	"encoding/json"

	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/common/logger"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/streadway/amqp"
)

type Consumer struct {
	channel *amqp.Channel
	queries *generated.Queries
}

type BookingAcceptedEvent struct {
	BookingID  string `json:"booking_id"`
	DriverID   string `json:"driver_id"`
	RideStatus string `json:"ride_status"`
}

func NewConsumer(channel *amqp.Channel, queries *generated.Queries) *Consumer {
	return &Consumer{
		channel: channel,
		queries: queries,
	}
}

func (c *Consumer) StartConsumingBookingAccepted(ctx context.Context) error {
	logger.Infof("Setting up booking.accepted queue...")

	// Declare the queue
	_, err := c.channel.QueueDeclare(
		"booking.accepted", // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		logger.Errorln("Failed to declare booking.accepted queue", err)
		return err
	}
	logger.Infof("booking.accepted queue declared")

	// Start consuming messages
	msgs, err := c.channel.Consume(
		"booking.accepted", // queue
		"booking-service",  // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	if err != nil {
		logger.Errorln("Failed to start consuming booking.accepted messages", err)
		return err
	}

	logger.Infof("Waiting for booking.accepted messages...")

	go func() {
		for {
			select {
			case <-ctx.Done():
				logger.Errorln("Consumer context cancelled, stopping...")
				return
			case msg := <-msgs:
				if len(msg.Body) == 0 {
					continue
				}

				logger.Infof("Received booking.accepted message: %s", string(msg.Body))

				var event BookingAcceptedEvent
				if err := json.Unmarshal(msg.Body, &event); err != nil {
					logger.Errorln("Failed to unmarshal booking.accepted event", err)
					continue
				}

				logger.Infof("Parsed booking.accepted event: %+v", event)

				// Update booking status and driver_id idempotently
				if err := c.handleBookingAccepted(ctx, event); err != nil {
					logger.Errorln("Failed to handle booking.accepted event", err)
				} else {
					logger.Infof("Successfully handled booking.accepted event for booking: %s", event.BookingID)
				}
			}
		}
	}()

	return nil
}

func (c *Consumer) handleBookingAccepted(ctx context.Context, event BookingAcceptedEvent) error {
	// Parse booking ID from string to UUID
	var bookingUUID pgtype.UUID
	err := bookingUUID.Scan(event.BookingID)
	if err != nil {
		logger.Errorln("Failed to parse booking ID", err)
		return err
	}

	// Parse driver ID from string to UUID
	var driverUUID pgtype.UUID
	err = driverUUID.Scan(event.DriverID)
	if err != nil {
		logger.Errorln("Failed to parse driver ID", err)
		return err
	}

	err = c.queries.UpdateBookingStatus(ctx, generated.UpdateBookingStatusParams{
		BookingID:  bookingUUID,
		RideStatus: event.RideStatus,
		DriverID:   driverUUID,
	})

	if err != nil {
		logger.Errorln("Failed to update booking status", err)
		return err
	}

	logger.Infof("Successfully updated booking %s to accepted with driver %s", event.BookingID, event.DriverID)
	return nil
}
