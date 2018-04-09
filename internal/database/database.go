package database

import (
	"context"
	"database/sql"
	"fmt"
	// This import is required because Package pq is a pure Go Postgres driver for the database/sql package.
	_ "github.com/lib/pq"
)

//New function to create new DB connection
func New(ctx context.Context, connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening db connection %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error pinging db: %v", err)
	}

	return db, nil
}
