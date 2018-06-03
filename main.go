package main

import (
	"github.com/kot13/vertigo/app"
	"github.com/kot13/vertigo/config"
	"github.com/kot13/vertigo/version"

	log "github.com/Sirupsen/logrus"
)

func main() {
	cfg := config.New()

	a := app.New(cfg)

	log.Infof("Starting the service. Commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release)

	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
