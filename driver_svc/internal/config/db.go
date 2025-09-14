package config

import (
	"context"

	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/db/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDb() (*pgxpool.Pool, *generated.Queries, error) {
	pool, err := pgxpool.New(context.Background(), LoadConfig().DBUrl)
	if err != nil {
		return nil, nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, nil, err
	}
	queries := generated.New(pool)
	return pool, queries, nil
}
