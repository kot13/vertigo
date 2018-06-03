package handlers

import (
	"net/http"

	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
)

func Get(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if ID == "" {
		renderer.Error("Param id is require", w)
		return
	}

	advert, err := container.GetDb().GetAdvertById(ID)
	if err != nil {
		renderer.Error("Error:"+err.Error(), w)
		return
	}

	renderer.Render(advert, w)
}
