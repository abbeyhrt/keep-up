package main

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/database"
	"github.com/abbeyhrt/keep-up-graphql/internal/pubapi/pubapisrv"
)

func main() {
	cfg, _ := config.New()
	ctx := context.Background()

	if cfg.Env == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	s, err := database.New(ctx, cfg.Postgres.String())
	if err != nil {
		log.Fatal(err)
	}

	srv, err := pubapisrv.New(ctx, cfg, s)
	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Listening at http://" + cfg.Addr)
	log.Fatal(srv.ListenAndServe())
}
