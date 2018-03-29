package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/database"
	"github.com/abbeyhrt/keep-up-graphql/internal/pubapi/pubapisrv"
)

func main() {
	cfg := config.New()

	if cfg.Env == "development" {
		_, err := database.New(cfg.Postgres.String())
		if err != nil {
			log.Fatal(err)
		}
	}

	srv, err := pubapisrv.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Listening at http://" + cfg.Addr)
	log.Fatal(srv.ListenAndServe())
}
