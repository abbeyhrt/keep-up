package pubapisrv_test

import (
	"context"
	"errors"
	"testing"

	"github.com/abbeyhrt/keep-up/graphql/internal/config"
	"github.com/abbeyhrt/keep-up/graphql/internal/models"
	"github.com/abbeyhrt/keep-up/graphql/internal/pubapi/pubapisrv"
)

type MockStore struct{}

func (s *MockStore) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}

func (s *MockStore) GetOrCreateUser(ctx context.Context, user *models.User) error {
	return errors.New("not implemented")
}

func (s *MockStore) GetUserByID(ctx context.Context, id string) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}

func (s *MockStore) CreateSession(ctx context.Context, userID string) (models.Session, error) {
	return models.Session{}, errors.New("not implemented")
}

func (s *MockStore) GetSessionByID(ctx context.Context, id string) (models.Session, error) {
	return models.Session{}, errors.New("not implemented")
}

func (s *MockStore) CreateHome(ctx context.Context, home models.Home, userID string) (models.Home, error) {
	return models.Home{}, errors.New("not implemented")
}

func (s *MockStore) GetHomeByID(ctx context.Context, homeID *string) (models.Home, error) {
	return models.Home{}, errors.New("not implemented")
}

func (s *MockStore) UpdateTask(ctx context.Context, task models.Task) (models.Task, error) {
	return models.Task{}, errors.New("not implemented")
}
func (s *MockStore) CreateTask(ctx context.Context, task models.Task, userID string) (models.Task, error) {
	return models.Task{}, errors.New("not implemented")
}

func (s *MockStore) GetTasksByUserID(ctx context.Context, userID string) ([]models.Task, error) {
	return []models.Task{}, errors.New("not implemented")
}

func (s *MockStore) GetTaskByID(ctx context.Context, id string) (models.Task, error) {
	return models.Task{}, errors.New("not implemented")
}

func (s *MockStore) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	return models.User{}, errors.New("not implemented")
}

func (s *MockStore) GetUsersByName(ctx context.Context, name string) ([]models.User, error) {
	return []models.User{}, errors.New("not implemented")
}

func (s *MockStore) UpdateHome(ctx context.Context, home models.Home) (models.Home, error) {
	return models.Home{}, errors.New("not implemented")
}

func TestNewServer(t *testing.T) {
	ctx := context.Background()
	addr := "0.0.0.0:3000"
	cfg := config.Config{Addr: addr}
	store := MockStore{}

	srv, err := pubapisrv.New(ctx, cfg, &store)
	if err != nil {
		t.Fatal(err)
	}

	if srv.Addr != addr {
		t.Errorf("Expected server address to be '%s', instead got '%s'", addr, srv.Addr)
	}
}
