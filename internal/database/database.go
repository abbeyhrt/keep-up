package database

import (
	"database/sql"
	"fmt"
	//i need it
	_ "github.com/lib/pq"
)

//New function to create new DB connection
func New(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening db connection %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db: %v", err)
	}

	return db, nil
}
