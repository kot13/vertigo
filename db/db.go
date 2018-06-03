package db

import (
	"github.com/kot13/vertigo/db/models"

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

func (db *DB) GetAdvertByHexId(HexID string) (advert models.AdvertData, err error) {
	_, err = db.QueryOne(&advert, `
		SELECT
			"id",
			"hex_id",
			"user_id",
			"properties",
			"created_at",
			"updated_at"
		FROM "advert_data"
		WHERE "hex_id" = ?
	`, HexID)

	return
}
