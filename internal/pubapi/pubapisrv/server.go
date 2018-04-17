package pubapisrv

import (
	"net/http"
	"time"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/handlers"
)

// New server instance
func New(cfg config.Config) (*http.Server, error) {
	handler := handlers.New(cfg)

	srv := http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &srv, nil
}
