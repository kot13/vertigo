package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-pg/pg"
)

type DB struct {
	*pg.DB
}

func New(dsn string) *DB {
	pgOpt, err := pg.ParseURL(dsn)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		pg.Connect(pgOpt),
	}
}
