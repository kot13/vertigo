package db

import (
	"github.com/kot13/vertigo/db/models"

	log "github.com/Sirupsen/logrus"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"strconv"
	"time"
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

func (db *DB) GetAdvertById(ID string) (advert models.AdvertData, err error) {
	_, err = db.QueryOne(&advert, `
		SELECT
			"id",
			"user_id",
			"properties",
			"created_at",
			"updated_at",
			"status"
		FROM "advert_data"
		WHERE "id" = ?
	`, ID)

	return
}

func (db *DB) SetStatus(ID string, status int8) (res orm.Result, err error) {
	iid, err := strconv.Atoi(ID)
	if err != nil {
		return
	}

	now := time.Now()

	advert := models.AdvertData{
		ID:        int64(iid),
		Status:    status,
		UpdatedAt: &now,
	}

	res, err = db.Model(&advert).Column("status", "updated_at").WherePK().Update()

	return
}

func (db *DB) SetProperties(ID string, props map[string]models.AdvertProperty) (res orm.Result, err error) {
	iid, err := strconv.Atoi(ID)
	if err != nil {
		return
	}

	now := time.Now()

	advert := models.AdvertData{
		ID:         int64(iid),
		Properties: props,
		UpdatedAt:  &now,
	}

	res, err = db.Model(&advert).Column("properties", "updated_at").WherePK().Update()

	return
}
