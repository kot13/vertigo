package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("seeding dictionary")
		_, err := db.Exec(`INSERT INTO dictionary (property, value, data) VALUES ('vendor', 1, 'Amati')`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("truncating dictionary")
		_, err := db.Exec(`TRUNCATE dictionary`)
		return err
	})
}
