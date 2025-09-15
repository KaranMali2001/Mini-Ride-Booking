package main

import (
	"context"
	"fmt"

	"github.com/KaranMali2001/mini-ride-booking/common/logger"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/config"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/handlers"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/mq"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/repo"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/service"
)

func bootstrapHandler(queries *generated.Queries) (*handlers.DriverHandler, error) {
	// Connect to RabbitMQ
	logger.Info("Connecting to RabbitMQ...")
	rabbitmqURL := config.LoadConfig().RabbitURL

	conn, ch := config.ConnectRabbit(rabbitmqURL)
	if conn == nil || ch == nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ")
	}

	consumer := mq.NewConsumer(ch, queries)
	ctx := context.Background()

	go consumer.StartConsumingBookingCreated(ctx)
	logger.Info("Consumer started successfully")

	producer := mq.NewProducer(ch)
	logger.Info("Producer created successfully")

	repo := repo.NewDriverRepo(queries)
	svc := service.NewDriverService(repo, producer)
	handler := handlers.NewDriverHandler(svc)
	logger.Info("Service layers built successfully")

	return handler, nil
}
