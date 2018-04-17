package main

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/database"
	"github.com/abbeyhrt/keep-up-graphql/internal/pubapi/pubapisrv"
)

func main() {
<<<<<<< HEAD
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
=======
	cfg := config.New()
	ctx := context.Background()

	if cfg.Env == "production" {
		log.SetFormatter(&log.JSONFormatter{})
>>>>>>> ce7db0c4387d73f1db896a0d51a42ec3d25b569b
	}

	db, err := database.New(ctx, cfg.Postgres.String())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	srv, err := pubapisrv.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Listening at http://" + cfg.Addr)
	log.Fatal(srv.ListenAndServe())
}
