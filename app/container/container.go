package container

import (
	"github.com/kot13/vertigo/config"
	"github.com/kot13/vertigo/db"

	log "github.com/Sirupsen/logrus"
)

type Container struct {
	cfg *config.Config
	db  *db.DB
}

var instance *Container

func New() *Container {
	instance = &Container{}
	return instance
}

func SetCfg(cfg *config.Config) {
	if instance.cfg != nil {
		log.Fatal("This property is already set")
	}
	instance.cfg = cfg
}

func SetDb(db *db.DB) {
	if instance.db != nil {
		log.Fatal("This property is already set")
	}
	instance.db = db
}

func GetCfg() *config.Config {
	if instance.cfg == nil {
		log.Fatal("This property is not initialized")
	}

	return instance.cfg
}

func GetDb() *db.DB {
	if instance.db == nil {
		log.Fatal("This property is not initialized")
	}

	return instance.db
}
