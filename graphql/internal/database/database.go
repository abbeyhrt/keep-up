package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/abbeyhrt/keep-up/graphql/internal/models"
	_ "github.com/lib/pq"
)

type DAL interface {
	CreateSession(ctx context.Context, userID string) (models.Session, error)
	GetSessionByID(ctx context.Context, id string) (models.Session, error)
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetOrCreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (models.User, error)
	CreateHome(ctx context.Context, home models.Home, userID string) (models.Home, error)
	GetHomeByID(ctx context.Context, homeID *string) (models.Home, error)
	CreateTask(ctx context.Context, task models.Task, userID string) (models.Task, error)
	GetTasksByUserID(ctx context.Context, userID string) ([]models.Task, error)
	GetTaskByID(ctx context.Context, id string) (models.Task, error)
	GetUsersByName(ctx context.Context, name string) ([]models.User, error)
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
