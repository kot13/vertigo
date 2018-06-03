package app

import (
	"net/http"
	"os"

	"github.com/kot13/vertigo/app/handlers"
	"github.com/kot13/vertigo/config"
	"github.com/kot13/vertigo/db"

	log "github.com/Sirupsen/logrus"
)

type App struct {
	cfg *config.Config
	db  *db.DB
}

func New(cfg *config.Config) *App {
	initLogger(cfg)

	return &App{
		cfg: cfg,
		db:  db.New(cfg.Db),
	}
}

func (app *App) Run() error {
	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static"+r.URL.Path)
	})

	http.HandleFunc("/docs", handlers.Docs)
	http.HandleFunc("/advert/get", handlers.Get)
	http.HandleFunc("/advert/search", handlers.Search)
	http.HandleFunc("/advert/create", handlers.Create)
	http.HandleFunc("/advert/update", handlers.Update)
	http.HandleFunc("/advert/publish", handlers.Publish)
	http.HandleFunc("/advert/un-publish", handlers.UnPublish)

	return http.ListenAndServe(":"+app.cfg.Port, nil)
}

func initLogger(cfg *config.Config) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Error parse log level: %s\n", err)
	} else {
		log.SetLevel(level)
	}
}
