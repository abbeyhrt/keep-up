package database

import (
	"context"
	"errors"
	"reflect"

	"github.com/abbeyhrt/keep-up/graphql/internal/models"
)

type MockStore struct {
	MockCreateSession    func(ctx context.Context, userID string) (models.Session, error)
	MockGetSessionByID   func(ctx context.Context, id string) (models.Session, error)
	MockCreateUser       func(ctx context.Context, user models.User) (models.User, error)
	MockGetOrCreateUser  func(ctx context.Context, user *models.User) error
	MockGetUserByID      func(ctx context.Context, id string) (models.User, error)
	MockCreateHome       func(ctx context.Context, home models.Home, userID string) (models.Home, error)
	MockGetHomeByID      func(ctx context.Context, homeID *string) (models.Home, error)
	MockCreateTask       func(ctx context.Context, task models.Task, userID string) (models.Task, error)
	MockGetTasksByUserID func(ctx context.Context, userID string) ([]models.Task, error)
	MockGetTaskByID      func(ctx context.Context, id string) (models.Task, error)
}

func (s *MockStore) Set(name string, method interface{}) {
	r := reflect.ValueOf(s)
	reflect.Indirect(r).FieldByName("Mock" + name).Set(reflect.ValueOf(method))
}

func NewMockStore() *MockStore {
	return &MockStore{
		MockCreateSession: func(ctx context.Context, userID string) (models.Session, error) {
			return models.Session{}, errors.New("not implemented")
		},
		MockGetSessionByID: func(ctx context.Context, id string) (models.Session, error) {
			return models.Session{}, errors.New("not implemented")
		},
		MockCreateUser: func(ctx context.Context, user models.User) (models.User, error) {
			return models.User{}, errors.New("not implemented")
		},
		MockGetOrCreateUser: func(ctx context.Context, user *models.User) error {
			return errors.New("not implemented")
		},
		MockGetUserByID: func(ctx context.Context, id string) (models.User, error) {
			return models.User{}, errors.New("not implemented")
		},
		MockCreateHome: func(ctx context.Context, home models.Home, userID string) (models.Home, error) {
			return models.Home{}, errors.New("not implemented")
		},
		MockGetHomeByID: func(ctx context.Context, homeID *string) (models.Home, error) {
			return models.Home{}, errors.New("not implemented")
		},
		MockCreateTask: func(ctx context.Context, task models.Task, userID string) (models.Task, error) {
			return models.Task{}, errors.New("not implemented")
		},
		MockGetTasksByUserID: func(ctx context.Context, userID string) ([]models.Task, error) {
			return []models.Task{}, errors.New("not implemented")
		},
		MockGetTaskByID: func(ctx context.Context, id string) (models.Task, error) {
			return models.Task{}, errors.New("not implemented")
		},
	}
}

func (s *MockStore) CreateSession(ctx context.Context, userID string) (models.Session, error) {
	return s.MockCreateSession(ctx, userID)
}

func (s *MockStore) GetSessionByID(ctx context.Context, id string) (models.Session, error) {
	return s.MockGetSessionByID(ctx, id)
}

func (s *MockStore) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return s.MockCreateUser(ctx, user)
}

func (s *MockStore) GetOrCreateUser(ctx context.Context, user *models.User) error {
	return s.MockGetOrCreateUser(ctx, user)
}

func (s *MockStore) GetUserByID(ctx context.Context, id string) (models.User, error) {
	return s.MockGetUserByID(ctx, id)
}

func (s *MockStore) CreateHome(ctx context.Context, home models.Home, userID string) (models.Home, error) {
	return s.MockCreateHome(ctx, home, userID)
}

func (s *MockStore) GetHomeByID(ctx context.Context, homeID *string) (models.Home, error) {
	return s.MockGetHomeByID(ctx, homeID)
}

func (s *MockStore) CreateTask(ctx context.Context, task models.Task, userID string) (models.Task, error) {
	return s.MockCreateTask(ctx, task, userID)
}

func (s *MockStore) GetTasksByUserID(ctx context.Context, userID string) ([]models.Task, error) {
	return s.MockGetTasksByUserID(ctx, userID)
}

func (s *MockStore) GetTaskByID(ctx context.Context, id string) (models.Task, error) {
	return s.MockGetTaskByID(ctx, id)
}
