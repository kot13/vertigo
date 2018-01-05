package main

import (
	"net/http"
	"os"

	"github.com/kot13/vertigo/handlers"
	"github.com/kot13/vertigo/version"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.FatalLevel)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		log.Fatal("Logger level is not set.")
	}

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal("Error parse log level: %s\n", err)
	} else {
		log.SetLevel(level)
	}

	log.Debugf("Starting the service. Commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release)

	http.HandleFunc("/health-check", handlers.HealthCheck)
	http.ListenAndServe(":"+port, nil)
}
