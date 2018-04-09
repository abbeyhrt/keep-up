package pubapisrv_test

import (
	"context"
	"testing"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/pubapi/pubapisrv"
)

func TestNewServer(t *testing.T) {
	ctx := context.Background()
	addr := "0.0.0.0:3000"
	cfg := config.Config{Addr: addr}

	srv, err := pubapisrv.New(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if srv.Addr != addr {
		t.Errorf("Expected server address to be '%s', instead got '%s'", addr, srv.Addr)
	}
}
