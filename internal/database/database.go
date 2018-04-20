package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/abbeyhrt/keep-up-graphql/internal/models"
	_ "github.com/lib/pq"
)

type DAL interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	FindOrCreateUser(ctx context.Context, user *models.User) error
}

func New(ctx context.Context, connStr string) (DAL, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening db connection %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error pinging db: %v", err)
	}

	return &SQLStore{db}, nil
}