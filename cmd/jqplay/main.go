package main

import (
	"context"

	"github.com/owenthereal/jqplay/config"
	"github.com/owenthereal/jqplay/jq"
	"github.com/owenthereal/jqplay/server"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	err = jq.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.WithFields(log.Fields{
		"version": jq.Version,
		"path":    jq.Path,
	}).Info("initialized jq")

	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.WithFields(log.Fields{
		"host": conf.Host,
		"port": conf.Port,
	}).Infof("Starting server at %s:%s", conf.Host, conf.Port)

	srv := server.New(conf)
	if err := srv.Start(context.Background()); err != nil {
		log.WithError(err).Fatal("error starting sever")
	}
}
