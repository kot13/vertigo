package renderer

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type ErrorResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Error(message string, w http.ResponseWriter) {
	log.Debug(message)

	e := ErrorResponse{
		Message: message,
	}

	Render(e, w)
}

func Render(res interface{}, w http.ResponseWriter) {
	response, err := json.Marshal(res)
	if err != nil {
		log.Errorf("Marshal error %+v", res)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
