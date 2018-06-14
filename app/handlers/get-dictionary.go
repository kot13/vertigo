package handlers

import (
	"net/http"

	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
	"github.com/kot13/vertigo/db/models"
)

func GetDictionary(w http.ResponseWriter, r *http.Request) {
	var dict []models.Dictionary

	q := container.GetDb().Model(&dict)

	property := r.URL.Query().Get("property")
	if property != "" {
		q.Where("dictionary.property = ?", property)
	}

	err := q.Select()
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	renderer.Render(dict, w)
}
