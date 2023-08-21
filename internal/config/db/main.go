package db

import (
	"database/sql"
	"fmt"

	"github.com/devigor/donna-markdown-service/internal/config/env"
)

func OpenConn() (*sql.DB, error) {
	connStr := env.GetEnv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error to connect in database", err.Error())
	} else {
		fmt.Println("Database Connected")
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
