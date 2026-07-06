package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"job-search/internal/auth"
)

var DB *pgx.Conn

func Connect() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}

	auth.SecretKey = []byte(os.Getenv("JWT_SECRET"))

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = pgx.Connect(
		context.Background(),
		connectionString,
	)

	if err != nil {
		return err
	}

	fmt.Println("PostgreSQL connected")

	return nil
}
