package handlers

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func Publish(w http.ResponseWriter, r *http.Request) {
	log.Debug("publish advert")
}
