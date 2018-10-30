package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/abbeyhrt/keep-up/graphql/internal/models"
	_ "github.com/lib/pq"
)

// DAL is the interface we are using to serve the store functions everywhere
type DAL interface {
	// Sessions
	CreateSession(ctx context.Context, userID string) (models.Session, error)
	GetSessionByID(ctx context.Context, id string) (models.Session, error)

	// Users
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetOrCreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (models.User, error)
	GetUsersByName(ctx context.Context, name string) ([]models.User, error)
	UpdateUser(ctx context.Context, user models.User) error

	// Home
	CreateHome(ctx context.Context, home models.Home, userID string) (models.Home, error)
	GetHomeByID(ctx context.Context, homeID *string) (models.Home, error)
	//InsertHomeID(ctx context.Context, userID string, homeID string) error

	// Tasks
	CreateTask(ctx context.Context, task models.Task, userID string) (models.Task, error)
	GetTasksByUserID(ctx context.Context, userID string) ([]models.Task, error)
	GetTaskByID(ctx context.Context, id string) (models.Task, error)
}

// New creates a new DAL
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
