package pubapisrv

import (
	"net/http"
	"time"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/handlers"
)

// New server instance
func New(cfg config.Config) (*http.Server, error) {
	// r := mux.NewRouter()
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello world"))
	// })

	handler := handlers.New()

	srv := http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &srv, nil
}
