package pubapisrv

import (
	"context"
	"net/http"
	"time"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/database"
	"github.com/abbeyhrt/keep-up-graphql/internal/handlers"
)

// New server instance
func New(
	ctx context.Context,
	cfg config.Config,
	store database.DAL,
) (*http.Server, error) {
	handler := handlers.New(ctx, cfg, store)

	srv := http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &srv, nil
}
