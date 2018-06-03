package handlers

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/kot13/vertigo/app/renderer"
)

func Search(w http.ResponseWriter, r *http.Request) {
	log.Debug("search advert")
	renderer.Error("Not implemented", w)
	return
}
