package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/abbeyhrt/keep-up/graphql/internal/models"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type DAL interface {

	//Sessions
	CreateSession(ctx context.Context, userID string) (models.Session, error)
	GetSessionByID(ctx context.Context, id string) (models.Session, error)

	//User
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetOrCreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (models.User, error)
	GetUsersByName(ctx context.Context, name string) ([]models.User, error)
	UpdateUser(ctx context.Context, user models.User) error

	//Home
	CreateHome(ctx context.Context, home models.Home, userID string) (models.Home, error)
	GetHomeByID(ctx context.Context, homeID *string) (models.Home, error)

	//Tasks
	CreateTask(ctx context.Context, task models.Task, userID string) (models.Task, error)
	GetTasksByUserID(ctx context.Context, userID string) ([]models.Task, error)
	GetTaskByID(ctx context.Context, id string) (models.Task, error)
}

func New(ctx context.Context, connStr string) (DAL, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening db connection %v", err)
	}

	err = db.PingContext(ctx)

	if err != nil {
		attempts := 0
		for attempts <= 6 {
			attempts++
			log.Infof("Retrying database connection, attempt #%d", attempts)
			if err := db.PingContext(ctx); err != nil {
				break
			}
			time.Sleep(10 * time.Second)
		}
		if err != nil {
			return nil, fmt.Errorf("error pinging db: %v", err)
		}
	}

	return &SQLStore{db}, nil
}
