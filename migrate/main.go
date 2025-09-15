package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("WARNING: .env file not found")
	}
	fmt.Println("Inside Migation", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("failed to open db", err)
		return
	}

	if err := goose.Up(db, "./migrations"); err != nil {
		fmt.Println("goose up failed", err)
		return
	}
	fmt.Println("Migration completed successfully")
}
