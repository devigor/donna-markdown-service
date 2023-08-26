package db

import (
	"context"
	"fmt"
	"os"

	"github.com/devigor/donna-notes-service/internal/config/env"
	"github.com/jackc/pgx/v5"
)

func OpenConn() (*pgx.Conn, error) {
	connStr := env.GetEnv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn, nil
}
