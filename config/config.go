package config

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

type Config struct {
	Port     string
	LogLevel string
	Db       string
}

func New() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		log.Fatal("Logger level is not set.")
	}

	db := os.Getenv("DATABASE")
	if db == "" {
		log.Fatal("Database is not set.")
	}

	return &Config{
		Port:     port,
		LogLevel: logLevel,
		Db:       db,
	}
}
