package db

import (
	"github.com/kot13/vertigo/db/models"

	log "github.com/Sirupsen/logrus"

	"fmt"
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

	dbc := pg.Connect(pgOpt)

	dbc.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Debugf("%s %s", time.Since(event.StartTime), query)
	})

	return &DB{
		dbc,
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

func (db *DB) SetProperties(ID string, props models.AdvertProperties) (res orm.Result, err error) {
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

func (db *DB) AddToIndex(ID string, props models.AdvertProperties) error {
	point := fmt.Sprintf("ST_GeomFromEWKT('SRID=4326;POINT(%s %s)')", *props.Lon, *props.Lat)

	_, err := db.Model((*models.AdvertIndex)(nil)).Exec(`
		INSERT INTO "advert_index" (id, price, vendor, point)
		VALUES (?, ?, ?, `+point+`)
	`, ID, *props.Price, *props.Vendor)

	return err
}

func (db *DB) RemoveFromIndex(ID string) error {
	iid, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}

	a := models.AdvertIndex{
		Id: int64(iid),
	}
	return db.Delete(&a)
}
