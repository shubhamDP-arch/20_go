package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Connect(ctx context.Context) *pgx.Conn {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	url := os.Getenv("databaseurl")
	fmt.Println(url)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		panic(err)
	}
	return conn
}
