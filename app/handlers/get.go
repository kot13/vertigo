package handlers

import (
	"net/http"

	"github.com/kot13/vertigo/app/renderer"
	"github.com/kot13/vertigo/db/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := string(r.URL.Query().Get("id"))

	if id == "" {
		renderer.Error("Param id not found", w)
	}

	renderer.Render(models.AdvertData{
		ID:    1,
		HexID: id,
	}, w)
}
