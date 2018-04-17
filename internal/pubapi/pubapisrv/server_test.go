package pubapisrv_test

import (
	"context"
	"errors"
	"testing"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/models"
	"github.com/abbeyhrt/keep-up-graphql/internal/pubapi/pubapisrv"
)

type MockStore struct{}

func (s *MockStore) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	return nil, errors.New("not implemented")
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
