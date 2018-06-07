package handlers

import (
	"net/http"

	"fmt"
	"github.com/kot13/vertigo/app/container"
	"github.com/kot13/vertigo/app/renderer"
	"github.com/kot13/vertigo/db/models"
)

var props = []SearchProperty{
	{
		Key:      "priceMin",
		Type:     "int64",
		Title:    "Цена от",
		Column:   "price",
		Operator: ">=",
	},
	{
		Key:      "priceMax",
		Type:     "int64",
		Title:    "Цена до",
		Column:   "price",
		Operator: "<=",
	},
}

func Search(w http.ResponseWriter, r *http.Request) {
	var adverts []models.AdvertIndex
	q := container.GetDb().Model(&adverts)

	values := r.URL.Query()
	for _, prop := range props {
		value := values.Get(prop.Key)
		if value != "" {
			q.Where(fmt.Sprintf("%s %s ?", prop.Column, prop.Operator), value)
		}
	}

	err := q.Limit(1).Select()
	if err != nil {
		renderer.Error(err.Error(), w)
		return
	}

	renderer.Render(adverts, w)
	return
}
