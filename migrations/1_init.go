package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating tables")
		_, err := db.Exec(`
			CREATE TABLE advert_data
			(
			  id         SERIAL                                 NOT NULL
				CONSTRAINT advert_data_pkey
				PRIMARY KEY,
			  user_id    BIGINT                                 NOT NULL,
			  properties JSONB,
			  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
			  updated_at TIMESTAMP WITH TIME ZONE,
			  status     SMALLINT                               NOT NULL
			);

			CREATE TABLE advert_index
			(
			  id     INTEGER NOT NULL
				CONSTRAINT advert_index_pkey
				PRIMARY KEY,
			  price  INTEGER NOT NULL,
			  vendor INTEGER NOT NULL
			);

			SELECT AddGeometryColumn('advert_index', 'point', 4326, 'POINT', 2);
			
			CREATE UNIQUE INDEX advert_index_id_uindex
			  ON advert_index (id);
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping tables")
		_, err := db.Exec(`DROP TABLE advert_data; DROP TABLE advert_index;`)
		return err
	})
}
