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

func (s *SQLStore) FindOrCreateUser(
	ctx context.Context,
	u *models.User,
) error {
	err := s.db.QueryRowContext(ctx, sqlGetUserByProvider, u.Provider, u.ProviderID).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.AvatarUrl,
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
		user.AvatarUrl,
		user.Provider,
		user.ProviderID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.AvatarUrl,
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

	sqlGetUserByProvider = `
	SELECT id, name, email, avatar_url, provider, provider_id, created_at, updated_at
	FROM users
	WHERE provider = $1 AND provider_id = $2
	`
)
