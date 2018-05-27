package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/abbeyhrt/keep-up/graphql/internal/models"
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

func (s *SQLStore) FindSessionByID(ctx context.Context, id string) (models.Session, error) {
	session := models.Session{}
	err := s.db.QueryRowContext(ctx, sqlGetSessionByID, id).Scan(
		&session.ID,
		&session.UserID,
	)
	return session, err
}

func (s *SQLStore) FindUserByID(ctx context.Context, id string) (models.User, error) {
	u := models.User{}
	err := s.db.QueryRowContext(ctx, sqlGetUserByID, id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.AvatarURL,
		&u.Provider,
		&u.ProviderID,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	return u, err
}

func (s *SQLStore) FindOrCreateUser(
	ctx context.Context,
	u *models.User,
) error {
	err := s.db.QueryRowContext(ctx, sqlGetUserByProvider, u.Provider, u.ProviderID).Scan(
		&u.ID,
		&u.Name,
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
	SELECT id, name, email, avatar_url, provider, provider_id, created_at, updated_at
	FROM users
	WHERE provider = $1 AND provider_id = $2
	`

	sqlGetUserByID = `
	SELECT id, name, email, avatar_url, provider, provider_id, created_at, updated_at
	FROM users
	WHERE id = $1
	`
)