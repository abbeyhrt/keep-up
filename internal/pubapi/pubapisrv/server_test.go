package pubapisrv_test

import (
	"testing"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/pubapi/pubapisrv"
)

func TestNewServer(t *testing.T) {
	addr := "0.0.0.0:3000"
	cfg := config.Config{Addr: addr}

	srv, err := pubapisrv.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if srv.Addr != addr {
		t.Errorf("Expected server address to be '%s', instead got '%s'", addr, srv.Addr)
	}
}
