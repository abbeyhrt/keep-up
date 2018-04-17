package pubapisrv

import (
	"context"
	"net/http"
	"time"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/handlers"
)

<<<<<<< HEAD
// New server instance
func New(cfg config.Config) (*http.Server, error) {
	handler := handlers.New(cfg)
=======
func New(ctx context.Context, cfg config.Config) (*http.Server, error) {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
>>>>>>> ce7db0c4387d73f1db896a0d51a42ec3d25b569b

	srv := http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &srv, nil
}
