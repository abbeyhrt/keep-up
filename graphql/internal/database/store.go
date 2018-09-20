package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/abbeyhrt/keep-up/graphql/internal/models"
	"github.com/opentracing/opentracing-go/log"
)

func NewStoreFromClient(db *sql.DB) *SQLStore {
	return &SQLStore{db}
}

type SQLStore struct {
	db *sql.DB
}

func (s *SQLStore) CreateSession(ctx context.Context, userID string) (models.Session, error) {
	session := models.Session{}
	err := s.db.QueryRowContext(
		ctx,
		sqlCreateSession,
		userID,
	).Scan(
		&session.ID,
		&session.UserID,
		&session.CreatedAt,
	)

	return session, err

}

func (s *SQLStore) GetSessionByID(ctx context.Context, id string) (models.Session, error) {
	session := models.Session{}
	err := s.db.QueryRowContext(ctx, sqlGetSessionByID, id).Scan(
		&session.ID,
		&session.UserID,
	)
	return session, err
}

func (s *SQLStore) GetUserByID(ctx context.Context, id string) (models.User, error) {
	u := models.User{}
	err := s.db.QueryRowContext(ctx, sqlGetUserByID, id).Scan(
		&u.ID,
		&u.Name,
		&u.HomeID,
		&u.Email,
		&u.AvatarURL,
		&u.Provider,
		&u.ProviderID,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	return u, err
}

func (s *SQLStore) GetOrCreateUser(
	ctx context.Context,
	u *models.User,
) error {
	err := s.db.QueryRowContext(ctx, sqlGetUserByProvider, u.Provider, u.ProviderID).Scan(
		&u.ID,
		&u.Name,
		&u.HomeID,
		&u.Email,
		&u.AvatarURL,
		&u.Provider,
		&u.ProviderID,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		user, err := s.CreateUser(ctx, *u)
		if err != nil {
			return err
		}

		*u = user
		return nil
	}

	return err
}

func (s *SQLStore) CreateUser(
	ctx context.Context,
	user models.User,
) (models.User, error) {
	err := s.db.QueryRowContext(
		ctx,
		sqlCreateUser,
		user.Name,
		user.Email,
		user.AvatarURL,
		user.Provider,
		user.ProviderID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.AvatarURL,
		&user.Provider,
		&user.ProviderID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, fmt.Errorf("error creating user: %v", err)
	}

	return user, nil
}

// CreateHome function that will be used in the handlers
func (s *SQLStore) CreateHome(ctx context.Context, home models.Home, userID string) (models.Home, error) {
	err := s.db.QueryRowContext(
		ctx,
		sqlCreateHome,
		home.Name,
		home.Description,
		home.AvatarURL,
	).Scan(
		&home.ID,
		&home.Name,
		&home.Description,
		&home.AvatarURL,
		&home.CreatedAt,
		&home.UpdatedAt,
	)
	if err != nil {
		return home, fmt.Errorf("error creating home: %v", err)
	}

	_, err = s.db.ExecContext(
		ctx,
		sqlInsertHomeID,
		home.ID,
		userID,
	)
	if err != nil {
		return home, fmt.Errorf("error creating home: %v", err)
	}

	return home, nil
}

//GetHomeByID used in handlers package
func (s *SQLStore) GetHomeByID(ctx context.Context, homeID *string) (models.Home, error) {
	h := models.Home{}
	err := s.db.QueryRowContext(ctx, sqlGetHomeByID, homeID).Scan(
		&h.ID,
		&h.Name,
		&h.Description,
	)

	return h, err
}

func (s *SQLStore) CreateTask(ctx context.Context, task models.Task, userID string) (models.Task, error) {
	err := s.db.QueryRowContext(
		ctx,
		sqlCreateTask,
		userID,
		task.Title,
		task.Description,
	).Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&task.Description,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return task, fmt.Errorf("error creating home: %v", err)
	}

	return task, nil

}

func (s *SQLStore) GetTasksByUserID(ctx context.Context, userID string) ([]models.Task, error) {
	rows, err := s.db.QueryContext(ctx, sqlGetTasksByUserID, userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		t := models.Task{}

		err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.Title,
			&t.Description,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (s *SQLStore) GetTaskByID(ctx context.Context, id string) (models.Task, error) {
	t := models.Task{}
	err := s.db.QueryRowContext(
		ctx,
		sqlGetTaskByID,
		id,
	).Scan(
		&t.ID,
		&t.UserID,
		&t.Title,
		&t.Description,
	)
	if err != nil {
		log.Error(err)
		fmt.Printf("this is the error from DB: %s", err)
		return t, err
	}

	return t, err
}

const (
	sqlCreateUser = `
	INSERT into users
	(name, email, avatar_url, provider, provider_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, name, email, avatar_url, provider, provider_id, created_at, updated_at
	`

	sqlCreateSession = `
	INSERT into sessions
	(user_id)
	VALUES ($1)
	RETURNING id, user_id, created_at
	`
	sqlGetSessionByID = `
	SELECT id, user_id
	FROM sessions
	WHERE id = $1`

	sqlGetUserByProvider = `
	SELECT id, name, home_id, email, avatar_url, provider, provider_id, created_at, updated_at
	FROM users
	WHERE provider = $1 AND provider_id = $2
	`

	sqlGetUserByID = `
	SELECT id, name, home_id, email, avatar_url, provider, provider_id, created_at, updated_at
	FROM users
	WHERE id = $1
	`

	sqlCreateHome = `
	INSERT into homes
	(name, description, avatar_url)
	VALUES ($1, $2, $3)
	RETURNING id, name, description, avatar_url, created_at, updated_at
	`

	sqlInsertHomeID = `
	UPDATE users
	SET home_id = $1
	WHERE id = $2
	`

	sqlGetHomeByID = `
	SELECT id, name, description
	FROM homes
	WHERE id = $1
	`

	sqlCreateTask = `
	INSERT into tasks
	(user_id, title, description)
	VALUES ($1, $2, $3)
	RETURNING id, user_id, title, description, created_at, updated_at
	`

	sqlGetTasksByUserID = `
	SELECT id, user_id, title, description, created_at, updated_at
	FROM tasks
	WHERE user_id = $1
	`

	sqlGetTaskByID = `
	SELECT id, user_id, title, description
	FROM tasks
	WHERE id = $1`
)
