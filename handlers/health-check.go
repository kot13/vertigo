package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kot13/vertigo/version"

	log "github.com/Sirupsen/logrus"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	info := struct {
		Status    string `json:"status"`
		BuildTime string `json:"buildTime"`
		Commit    string `json:"commit"`
		Release   string `json:"release"`
	}{
		"ok", version.BuildTime, version.Commit, version.Release,
	}

	renderResponse(info, w)
}

func renderResponse(res interface{}, w http.ResponseWriter) {
	response, err := json.Marshal(res)
	if err != nil {
		log.Printf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
