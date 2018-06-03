package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-pg/pg"
	"os"
)

var conn *pg.DB

// InitDB init connect to DB
func InitDB() {
	cfg := os.Getenv("DATABASE")
	pgOpt, err := pg.ParseURL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	conn = pg.Connect(pgOpt)
}
