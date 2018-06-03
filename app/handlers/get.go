package handlers

import (
	"net/http"

	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
)

func Get(w http.ResponseWriter, r *http.Request) {
	hexID := r.URL.Query().Get("hex_id")

	if hexID == "" {
		renderer.Error("Param hex_id is require", w)
		return
	}

	advert, err := container.GetDb().GetAdvertByHexId(hexID)
	if err != nil {
		renderer.Error("Error:"+err.Error(), w)
		return
	}

	renderer.Render(advert, w)
}
