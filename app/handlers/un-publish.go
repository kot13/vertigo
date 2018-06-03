package handlers

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func UnPublish(w http.ResponseWriter, r *http.Request) {
	log.Debug("un publish advert")
}
