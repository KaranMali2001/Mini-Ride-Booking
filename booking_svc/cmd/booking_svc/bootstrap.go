package main

import (
	"context"

	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/config"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/handlers"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/mq"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/repo"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/service"
	"github.com/KaranMali2001/mini-ride-booking/common/logger"
)

func bootstrapHandler(queries *generated.Queries) (*handlers.BookingHandler, error) {

	logger.Info("Connecting to RabbitMQ...")
	rabbitmqURL := config.LoadConfig().RabbitURL

	_, ch := config.ConnectRabbit(rabbitmqURL)

	logger.Info("RabbitMQ connected successfully")

	logger.Info("Starting booking.accepted consumer...")
	consumer := mq.NewConsumer(ch, queries)
	ctx := context.Background()
	go consumer.StartConsumingBookingAccepted(ctx)
	logger.Info("Consumer started successfully")

	// Create producer
	logger.Info("Creating message producer...")
	producer := mq.NewProducer(ch)
	logger.Info("Producer created successfully")

	logger.Info("Building service layers...")
	repo := repo.NewBookingRepo(queries)
	svc := service.NewBookingService(repo, producer)
	handler := handlers.NewBookingHandler(svc)
	logger.Info("Service layers built successfully")

	return handler, nil
}
