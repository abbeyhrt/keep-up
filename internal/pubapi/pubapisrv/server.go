package pubapisrv

import (
	"net/http"
	"time"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/gorilla/mux"
)

func New(cfg config.Config) (*http.Server, error) {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!!"))
	})

	srv := http.Server{
		Addr:         cfg.Addr,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &srv, nil
}
