package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/keep-up-graphql/internal/config"
	"github.com/keep-up-graphql/internal/pubapi/pubapisrv"
)

func main() {
	cfg := config.New()

	if cfg.Env == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	srv, err := pubapisrv.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Listening at http://" + cfg.Addr)
	log.Fatal(srv.ListenAndServe())
}
