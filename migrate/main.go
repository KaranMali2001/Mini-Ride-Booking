package main

import (
	"database/sql"
	"os"

	"github.com/KaranMali2001/mini-ride-booking/common/logger"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Errorln("failed to open db: %v", err)
		return
	}

	if err := goose.Up(db, "./migrations"); err != nil {
		logger.Errorln("goose up failed: %v", err)
		return
	}
	logger.Info("Migration completed successfully")
}
