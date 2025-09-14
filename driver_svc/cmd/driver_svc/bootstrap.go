package main

import (
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/handlers"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/repo"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/service"
)

func bootstrapHandler(queries *generated.Queries) (*handlers.DriverHandler, error) {

	repo := repo.NewDriverRepo(queries)
	svc := service.NewDriverService(repo)
	handler := handlers.NewDriverHandler(svc)

	return handler, nil
}